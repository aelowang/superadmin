<script setup lang="tsx">
import { GetSiteTopupList } from '@/service/api';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import SiteTopupSearch from './modules/site-topup-search.vue';

const appStore = useAppStore();

const {
  columns,
  columnChecks,
  data,
  getData,
  getDataByPage,
  loading,
  mobilePagination,
  searchParams,
  resetSearchParams
} = useTable({
  apiFn: GetSiteTopupList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    siteUuid: null,
    topupStatus: null,
    siteName: null
  },
  columns: () => [
    {
      key: 'index',
      title: $t('common.index'),
      align: 'center',
      width: 64
    },
    {
      key: 'siteName',
      title: $t('page.site.siteName'),
      align: 'center',
      minWidth: 120
    },
    {
      key: 'score',
      title: $t('page.site.topup.score'),
      align: 'center',
      width: 120
    },
    {
      key: 'priceCurrency',
      title: $t('page.site.priceCurrency'),
      align: 'center',
      width: 100
    },
    {
      key: 'topupMethod',
      title: $t('page.site.topup.topupMethod'),
      align: 'center',
      width: 100
    },
    {
      key: 'topupStatus',
      title: $t('page.site.topup.topupStatus'),
      align: 'center',
      width: 100
    },
    {
      key: 'remark',
      title: $t('page.site.remark'),
      align: 'center',
      minWidth: 150,
      ellipsis: { tooltip: true }
    },
    {
      key: 'createTime',
      title: $t('page.site.topup.operateTime'),
      align: 'center',
      width: 170
    }
  ]
});

</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <SiteTopupSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard :title="$t('page.site.topup.title')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">

      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1100"
        :loading="loading"
        remote
        :row-key="row => row.uuid"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
