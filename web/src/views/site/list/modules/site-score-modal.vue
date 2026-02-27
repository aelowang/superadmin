<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { useLoading } from '@sa/hooks';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { DeductionSite, TopupSite } from '@/service/api';
import { $t } from '@/locales';

defineOptions({
  name: 'SiteScoreModal'
});

type ScoreType = 'topup' | 'deduction';

interface Props {
  type: ScoreType;
  siteUuid: string;
  siteName: string;
  remainingScore: string;
}

const props = defineProps<Props>();

interface Emits {
  (e: 'submitted'): void;
}

const emit = defineEmits<Emits>();

const visible = defineModel<boolean>('visible', {
  default: false
});

const { formRef, validate, restoreValidation } = useNaiveForm();
const { defaultRequiredRule } = useFormRules();
const { loading: confirmLoading, startLoading: confirmStartLoading, endLoading: confirmEndLoading } = useLoading();

const title = computed(() => {
  return props.type === 'topup' ? $t('page.site.scoreModal.topupTitle') : $t('page.site.scoreModal.deductionTitle');
});

type Model = {
  score: string;
  remark: string;
};

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    score: '',
    remark: ''
  };
}

const rules = {
  score: defaultRequiredRule
};

function closeModal() {
  visible.value = false;
}

async function handleSubmit() {
  await validate();

  const scoreNum = Number(model.score);
  if (Number.isNaN(scoreNum) || scoreNum <= 0) {
    window.$message?.error(`${$t('page.site.scoreModal.score')} > 0`);
    return;
  }

  if (props.type === 'deduction') {
    const remaining = Number(props.remainingScore);
    if (scoreNum > remaining) {
      window.$message?.error(`${$t('page.site.scoreModal.currentScore')}: ${props.remainingScore}`);
      return;
    }
  }

  confirmStartLoading();

  const reqData = {
    uuid: props.siteUuid,
    score: model.score,
    remark: model.remark
  };

  const { error } = props.type === 'topup' ? await TopupSite(reqData) : await DeductionSite(reqData);

  if (!error) {
    window.$message?.success(props.type === 'topup' ? $t('common.addSuccess') : $t('common.addSuccess'));
    closeModal();
    emit('submitted');
  }
  confirmEndLoading();
}

watch(visible, () => {
  if (visible.value) {
    Object.assign(model, createDefaultModel());
    restoreValidation();
  }
});
</script>

<template>
  <NModal v-model:show="visible" :title="title" preset="card" class="w-500px">
    <NForm ref="formRef" :model="model" :rules="rules" label-placement="left" :label-width="120">
      <NFormItem :label="$t('page.site.siteName')">
        <NInput :value="props.siteName" disabled />
      </NFormItem>
      <NFormItem :label="$t('page.site.scoreModal.currentScore')">
        <NInput :value="props.remainingScore" disabled />
      </NFormItem>
      <NFormItem :label="$t('page.site.scoreModal.score')" path="score">
        <NInput v-model:value="model.score" :placeholder="$t('page.site.scoreModal.scorePlaceholder')" />
      </NFormItem>
      <NFormItem :label="$t('page.site.remark')" path="remark">
        <NInput
          v-model:value="model.remark"
          type="textarea"
          :placeholder="$t('page.site.scoreModal.remarkPlaceholder')"
        />
      </NFormItem>
    </NForm>
    <template #footer>
      <NSpace :size="16" justify="end">
        <NButton @click="closeModal">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" :loading="confirmLoading" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>
</template>

<style scoped></style>
