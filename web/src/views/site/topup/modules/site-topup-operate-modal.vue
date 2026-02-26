<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { useLoading } from '@sa/hooks';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { AddSiteTopup, GetAllSites } from '@/service/api';
import { $t } from '@/locales';

defineOptions({
  name: 'SiteTopupOperateModal'
});

interface Props {
  operateType: NaiveUI.TableOperateType;
  rowData?: Api.Site.SiteTopup | null;
}

defineProps<Props>();

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

type Model = {
  siteUuid: string;
  score: string;
  topupMethod: string;
  topupStatus: string;
  remark: string;
};

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    siteUuid: '',
    score: '',
    topupMethod: '',
    topupStatus: '',
    remark: ''
  };
}

const rules = {
  siteUuid: defaultRequiredRule,
  score: defaultRequiredRule,
  topupMethod: defaultRequiredRule
};

const siteOptions = ref<CommonType.Option<string>[]>([]);

async function getSiteOptions() {
  const { error, data } = await GetAllSites();
  if (!error) {
    siteOptions.value = data.map(item => ({
      label: item.siteName,
      value: item.uuid
    }));
  }
}

function closeModal() {
  visible.value = false;
}

async function handleSubmit() {
  await validate();
  confirmStartLoading();
  const { error } = await AddSiteTopup({ ...model });
  if (!error) {
    window.$message?.success($t('common.addSuccess'));
    closeModal();
    emit('submitted');
  }
  confirmEndLoading();
}

watch(visible, () => {
  if (visible.value) {
    Object.assign(model, createDefaultModel());
    restoreValidation();
    getSiteOptions();
  }
});
</script>

<template>
  <NModal v-model:show="visible" :title="$t('page.site.topup.addTopup')" preset="card" class="w-600px">
    <NForm ref="formRef" :model="model" :rules="rules" label-placement="left" :label-width="100">
      <NFormItem :label="$t('page.site.siteName')" path="siteUuid">
        <NSelect
          v-model:value="model.siteUuid"
          :options="siteOptions"
          :placeholder="$t('page.site.topup.form.selectSite')"
          filterable
        />
      </NFormItem>
      <NFormItem :label="$t('page.site.topup.score')" path="score">
        <NInput v-model:value="model.score" :placeholder="$t('page.site.topup.form.score')" />
      </NFormItem>
      <NFormItem :label="$t('page.site.topup.topupMethod')" path="topupMethod">
        <NInput v-model:value="model.topupMethod" :placeholder="$t('page.site.topup.form.topupMethod')" />
      </NFormItem>
      <NFormItem :label="$t('page.site.topup.topupStatus')" path="topupStatus">
        <NInput v-model:value="model.topupStatus" :placeholder="$t('page.site.topup.form.topupStatus')" />
      </NFormItem>
      <NFormItem :label="$t('page.site.remark')" path="remark">
        <NInput v-model:value="model.remark" type="textarea" :placeholder="$t('page.site.form.remark')" />
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
