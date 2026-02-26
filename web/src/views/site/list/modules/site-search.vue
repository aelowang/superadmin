<script setup lang="ts">
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { enableStatusOptions } from '@/constants/business';
import { translateOptions } from '@/utils/common';

defineOptions({
  name: 'SiteSearch'
});

interface Emits {
  (e: 'reset'): void;
  (e: 'search'): void;
}

const emit = defineEmits<Emits>();

const { formRef, validate, restoreValidation } = useNaiveForm();

const model = defineModel<Api.Site.SiteSearchParams>('model', { required: true });

async function reset() {
  await restoreValidation();
  emit('reset');
}

async function search() {
  await validate();
  emit('search');
}
</script>

<template>
  <NCard :bordered="false" size="small" class="card-wrapper">
    <NCollapse>
      <NCollapseItem :title="$t('common.search')" name="site-search">
        <NForm ref="formRef" :model="model" label-placement="left" :label-width="100">
          <NGrid responsive="screen" item-responsive>
            <NFormItemGi span="24 s:12 m:6" :label="$t('page.site.siteName')" path="siteName" class="pr-24px">
              <NInput v-model:value="model.siteName" :placeholder="$t('page.site.form.siteName')" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('page.site.priceCurrency')" path="priceCurrency" class="pr-24px">
              <NInput v-model:value="model.priceCurrency" :placeholder="$t('page.site.form.priceCurrency')" />
            </NFormItemGi>
            <NFormItemGi span="24 s:12 m:6" :label="$t('page.site.status')" path="status" class="pr-24px">
              <NSelect
                v-model:value="model.status"
                :placeholder="$t('page.site.form.status')"
                :options="translateOptions(enableStatusOptions)"
                clearable
              />
            </NFormItemGi>
            <NFormItemGi span="24 m:6" class="pr-24px">
              <NSpace class="w-full" justify="end">
                <NButton @click="reset">
                  <template #icon>
                    <icon-ic-round-refresh class="text-icon" />
                  </template>
                  {{ $t('common.reset') }}
                </NButton>
                <NButton type="primary" ghost @click="search">
                  <template #icon>
                    <icon-ic-round-search class="text-icon" />
                  </template>
                  {{ $t('common.search') }}
                </NButton>
              </NSpace>
            </NFormItemGi>
          </NGrid>
        </NForm>
      </NCollapseItem>
    </NCollapse>
  </NCard>
</template>

<style scoped></style>
