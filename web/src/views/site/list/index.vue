<script setup lang="tsx">
import { reactive } from 'vue';
import { NButton, NPopconfirm, NTag } from 'naive-ui';
import { DeleteSite, GetSiteList } from '@/service/api';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { enableStatusRecord } from '@/constants/business';
import { useTable, useTableOperate } from '@/hooks/common/table';
import SiteOperateModal from './modules/site-operate-modal.vue';
import SiteSearch from './modules/site-search.vue';

const appStore = useAppStore();

type LoadingStatus = Record<string, boolean>;
const deleteLoadingStatus = reactive<LoadingStatus>({});

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
  apiFn: GetSiteList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    siteName: null,
    priceCurrency: null,
    status: null
  },
  columns: () => [
    {
      type: 'selection',
      align: 'center',
      width: 48
    },
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
      key: 'priceCurrency',
      title: $t('page.site.priceCurrency'),
      align: 'center',
      width: 100
    },
    {
      key: 'remainingScore',
      title: $t('page.site.remainingScore'),
      align: 'center',
      width: 120
    },
    {
      key: 'totalTopup',
      title: $t('page.site.totalTopup'),
      align: 'center',
      width: 120
    },
    {
      key: 'adminUrl',
      title: $t('page.site.adminUrl'),
      align: 'center',
      minWidth: 150,
      ellipsis: { tooltip: true }
    },
    {
      key: 'contactName',
      title: $t('page.site.contactName'),
      align: 'center',
      width: 100
    },
    {
      key: 'contactPhone',
      title: $t('page.site.contactPhone'),
      align: 'center',
      width: 120
    },
    {
      key: 'status',
      title: $t('page.site.status'),
      align: 'center',
      width: 80,
      render: row => {
        if (row.status === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };
        const label = $t(enableStatusRecord[row.status]);
        return <NTag type={tagMap[row.status]}>{label}</NTag>;
      }
    },
    {
      key: 'operate',
      title: $t('common.operate'),
      align: 'center',
      width: 130,
      render: row => (
        <div class="flex-center gap-8px">
          <NButton type="primary" ghost size="small" onClick={() => edit(row.uuid)}>
            {$t('common.edit')}
          </NButton>
          <NPopconfirm onPositiveClick={() => handleDelete(row.uuid)}>
            {{
              default: () => $t('common.confirmDelete'),
              trigger: () => (
                <NButton loading={deleteLoadingStatus[row.uuid]} type="error" ghost size="small">
                  {$t('common.delete')}
                </NButton>
              )
            }}
          </NPopconfirm>
        </div>
      )
    }
  ]
});

const {
  drawerVisible,
  operateType,
  editingData,
  handleAdd,
  handleEdit,
  checkedRowKeys,
  onBatchDeleted,
  onDeleted
} = useTableOperate(data, getData);

async function handleBatchDelete() {
  const uuids: string[] = checkedRowKeys.value.map(uuid => uuid);
  uuids.forEach(uuid => {
    deleteLoadingStatus[uuid] = true;
  });
  const { error } = await DeleteSite(uuids);
  uuids.forEach(uuid => {
    deleteLoadingStatus[uuid] = false;
  });
  if (error) return;
  onBatchDeleted();
}

async function handleDelete(uuid: string) {
  const uuids: string[] = [uuid];
  deleteLoadingStatus[uuid] = true;
  const { error } = await DeleteSite(uuids);
  deleteLoadingStatus[uuid] = false;
  if (error) return;
  onDeleted();
}

function edit(uuid: string) {
  handleEdit(uuid);
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <SiteSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard :title="$t('page.site.title')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :disabled-delete="checkedRowKeys.length === 0"
          :loading="loading"
          @add="handleAdd"
          @delete="handleBatchDelete"
          @refresh="getData"
        />
      </template>
      <NDataTable
        v-model:checked-row-keys="checkedRowKeys"
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1400"
        :loading="loading"
        remote
        :row-key="row => row.uuid"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
      <SiteOperateModal
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getDataByPage"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
