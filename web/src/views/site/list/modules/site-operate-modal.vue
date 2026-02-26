<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { useLoading } from '@sa/hooks';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { AddSite, EditSite } from '@/service/api';
import { $t } from '@/locales';
import { enableStatusOptions } from '@/constants/business';

defineOptions({
  name: 'SiteOperateModal'
});

interface Props {
  operateType: NaiveUI.TableOperateType;
  rowData?: Api.Site.Site | null;
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
  const titles: Record<NaiveUI.TableOperateType, string> = {
    add: $t('page.site.addSite'),
    edit: $t('page.site.editSite')
  };
  return titles[props.operateType];
});

type Model = Pick<
  Api.Site.AddSiteRequest,
  | 'siteName'
  | 'siteLogo'
  | 'priceCurrency'
  | 'scoreStatType'
  | 'status'
  | 'remark'
  | 'dbHost'
  | 'dbUsername'
  | 'dbPassword'
  | 'dbPort'
  | 'dbName'
  | 'jwtSecret'
  | 'adminUrl'
  | 'adminUsername'
  | 'contactName'
  | 'contactPhone'
  | 'contactEmail'
>;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    siteName: '',
    siteLogo: '',
    priceCurrency: '',
    scoreStatType: '',
    status: '1',
    remark: '',
    dbHost: '',
    dbUsername: '',
    dbPassword: '',
    dbPort: 3306,
    dbName: '',
    jwtSecret: '',
    adminUrl: '',
    adminUsername: '',
    contactName: '',
    contactPhone: '',
    contactEmail: ''
  };
}

type RuleKey = Extract<keyof Model, 'siteName' | 'status' | 'dbHost' | 'dbUsername' | 'dbPassword' | 'dbName'>;

const rules: Record<RuleKey, App.Global.FormRule> = {
  siteName: defaultRequiredRule,
  status: defaultRequiredRule,
  dbHost: defaultRequiredRule,
  dbUsername: defaultRequiredRule,
  dbPassword: defaultRequiredRule,
  dbName: defaultRequiredRule
};

function handleInitModel() {
  Object.assign(model, createDefaultModel());
  if (props.operateType === 'edit' && props.rowData) {
    Object.assign(model, props.rowData);
  }
}

function closeModal() {
  visible.value = false;
}

async function handleSubmit() {
  await validate();

  if (props.operateType === 'add') {
    confirmStartLoading();
    const { error } = await AddSite({ ...model });
    if (!error) {
      window.$message?.success($t('common.addSuccess'));
      closeModal();
      emit('submitted');
    }
    confirmEndLoading();
  } else if (props.operateType === 'edit') {
    confirmStartLoading();
    const { error } = await EditSite({
      uuid: props.rowData?.uuid,
      ...model
    });
    if (!error) {
      window.$message?.success($t('common.updateSuccess'));
      closeModal();
      emit('submitted');
    }
    confirmEndLoading();
  }
}

watch(visible, () => {
  if (visible.value) {
    handleInitModel();
    restoreValidation();
  }
});
</script>

<template>
  <NModal v-model:show="visible" :title="title" preset="card" class="w-700px">
    <NScrollbar class="h-480px pr-20px">
      <NForm ref="formRef" :model="model" :rules="rules" label-placement="left" :label-width="120">
        <NDivider title-placement="left">{{ $t('page.site.basicInfo') }}</NDivider>
        <NFormItem :label="$t('page.site.siteName')" path="siteName">
          <NInput v-model:value="model.siteName" :placeholder="$t('page.site.form.siteName')" />
        </NFormItem>
        <!-- <NFormItem :label="$t('page.site.siteLogo')" path="siteLogo">
          <NInput v-model:value="model.siteLogo" :placeholder="$t('page.site.form.siteLogo')" />
        </NFormItem> -->
        <NFormItem :label="$t('page.site.priceCurrency')" path="priceCurrency">
          <NInput v-model:value="model.priceCurrency" :placeholder="$t('page.site.form.priceCurrency')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.status')" path="status">
          <NRadioGroup v-model:value="model.status">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.site.remark')" path="remark">
          <NInput v-model:value="model.remark" type="textarea" :placeholder="$t('page.site.form.remark')" />
        </NFormItem>

        <NDivider title-placement="left">{{ $t('page.site.dbInfo') }}</NDivider>
        <NFormItem :label="$t('page.site.dbHost')" path="dbHost">
          <NInput v-model:value="model.dbHost" :placeholder="$t('page.site.form.dbHost')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.dbPort')" path="dbPort">
          <NInputNumber v-model:value="model.dbPort" :placeholder="$t('page.site.form.dbPort')" class="w-full" />
        </NFormItem>
        <NFormItem :label="$t('page.site.dbName')" path="dbName">
          <NInput v-model:value="model.dbName" :placeholder="$t('page.site.form.dbName')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.dbUsername')" path="dbUsername">
          <NInput v-model:value="model.dbUsername" :placeholder="$t('page.site.form.dbUsername')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.dbPassword')" path="dbPassword">
          <NInput
            v-model:value="model.dbPassword"
            type="password"
            show-password-on="click"
            :placeholder="$t('page.site.form.dbPassword')"
          />
        </NFormItem>

        <NDivider title-placement="left">{{ $t('page.site.adminInfo') }}</NDivider>
        <NFormItem :label="$t('page.site.jwtSecret')" path="jwtSecret">
          <NInput v-model:value="model.jwtSecret" :placeholder="$t('page.site.form.jwtSecret')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.adminUrl')" path="adminUrl">
          <NInput v-model:value="model.adminUrl" :placeholder="$t('page.site.form.adminUrl')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.adminUsername')" path="adminUsername">
          <NInput v-model:value="model.adminUsername" :placeholder="$t('page.site.form.adminUsername')" />
        </NFormItem>

        <NDivider title-placement="left">{{ $t('page.site.contactInfo') }}</NDivider>
        <NFormItem :label="$t('page.site.contactName')" path="contactName">
          <NInput v-model:value="model.contactName" :placeholder="$t('page.site.form.contactName')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.contactPhone')" path="contactPhone">
          <NInput v-model:value="model.contactPhone" :placeholder="$t('page.site.form.contactPhone')" />
        </NFormItem>
        <NFormItem :label="$t('page.site.contactEmail')" path="contactEmail">
          <NInput v-model:value="model.contactEmail" :placeholder="$t('page.site.form.contactEmail')" />
        </NFormItem>
      </NForm>
    </NScrollbar>
    <template #footer>
      <NSpace :size="16" justify="end">
        <NButton @click="closeModal">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" :loading="confirmLoading" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>
</template>

<style scoped></style>
