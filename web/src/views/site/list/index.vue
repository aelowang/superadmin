<script setup lang="tsx">
import { reactive, ref } from 'vue';
import { NButton, NPopconfirm, NTag } from 'naive-ui';
import { DeleteSite, GetSiteList, SyncRemoteSiteScore } from '@/service/api';
import { $t } from '@/locales';
import { useAuth } from '@/hooks/business/auth';
import { useAppStore } from '@/store/modules/app';
import { enableStatusRecord } from '@/constants/business';
import { useTable, useTableOperate } from '@/hooks/common/table';
import SiteOperateModal from './modules/site-operate-modal.vue';
import SiteScoreModal from './modules/site-score-modal.vue';
import SiteSearch from './modules/site-search.vue';

const appStore = useAppStore();
const { hasAuth } = useAuth();

type LoadingStatus = Record<string, boolean>;
const deleteLoadingStatus = reactive<LoadingStatus>({});
const syncLoadingStatus = reactive<LoadingStatus>({});

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
      width: 360,
      render: row => (
        <div class="flex-center gap-8px">
          {hasAuth('v1:manage:site:topup') && (
            <NButton type="success" ghost size="small" onClick={() => openScoreModal('topup', row)}>
              {$t('page.site.quickTopup')}
            </NButton>
          )}
          {hasAuth('v1:manage:site:deduction') && (
            <NButton type="warning" ghost size="small" onClick={() => openScoreModal('deduction', row)}>
              {$t('page.site.quickDeduction')}
            </NButton>
          )}
          {hasAuth('v1:manage:site:syncRemoteScore') && (
            <NButton
              type="info"
              ghost
              size="small"
              loading={syncLoadingStatus[row.uuid]}
              onClick={() => handleSyncRemoteScore(row.uuid)}
            >
              {$t('page.site.syncRemoteScore')}
            </NButton>
          )}
          {hasAuth('v1:manage:site:edit') && (
            <NButton type="primary" ghost size="small" onClick={() => edit(row.uuid)}>
              {$t('common.edit')}
            </NButton>
          )}
          {hasAuth('v1:manage:site:delete') && (
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
          )}
        </div>
      )
    }
  ]
});

const { drawerVisible, operateType, editingData, handleAdd, handleEdit, checkedRowKeys, onBatchDeleted, onDeleted } =
  useTableOperate(data, getData);

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

async function handleSyncRemoteScore(uuid: string) {
  syncLoadingStatus[uuid] = true;
  const { error } = await SyncRemoteSiteScore({ uuid });
  syncLoadingStatus[uuid] = false;
  if (error) return;
  getDataByPage();
}

function edit(uuid: string) {
  handleEdit(uuid);
}

const scoreModalVisible = ref(false);
const scoreModalType = ref<'topup' | 'deduction'>('topup');
const scoreModalSiteUuid = ref('');
const scoreModalSiteName = ref('');
const scoreModalRemainingScore = ref('');

function openScoreModal(type: 'topup' | 'deduction', row: Api.Site.Site) {
  scoreModalType.value = type;
  scoreModalSiteUuid.value = row.uuid;
  scoreModalSiteName.value = row.siteName;
  scoreModalRemainingScore.value = row.remainingScore;
  scoreModalVisible.value = true;
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <SiteSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard :title="$t('page.site.title')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          add-auth="v1:manage:site:add"
          delete-auth="v1:manage:site:delete"
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
        :scroll-x="1600"
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
      <SiteScoreModal
        v-model:visible="scoreModalVisible"
        :type="scoreModalType"
        :site-uuid="scoreModalSiteUuid"
        :site-name="scoreModalSiteName"
        :remaining-score="scoreModalRemainingScore"
        @submitted="getDataByPage"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
