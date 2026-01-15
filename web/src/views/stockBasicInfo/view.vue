<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="股票基础信息表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                股票代码
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                ii
              </template>
              {{ formValue.ii }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                ei
              </template>
              {{ formValue.ei }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易所名称
              </template>
              {{ formValue.exchange }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                name
              </template>
              {{ formValue.name }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                short_name
              </template>
              {{ formValue.shortName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                en_name
              </template>
              {{ formValue.enName }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                上市日期
              </template>
              {{ formValue.od }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                data_update_date
              </template>
              {{ formValue.dataUpdateDate }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                pc
              </template>
              {{ formValue.pc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                up
              </template>
              {{ formValue.up }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                dp
              </template>
              {{ formValue.dp }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                pk
              </template>
              {{ formValue.pk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                fv
              </template>
              {{ formValue.fv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                tv
              </template>
              {{ formValue.tv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                流通股比例 (%)
              </template>
              {{ formValue.floatRatio }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                is
              </template>
              {{ formValue.is }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易状态描述
              </template>
              {{ formValue.tradingStatus }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                industry
              </template>
              {{ formValue.industry }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                sector
              </template>
              {{ formValue.sector }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                market_type
              </template>
              {{ formValue.marketType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                data_source
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                is_active
              </template>
              {{ formValue.isActive }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                version
              </template>
              {{ formValue.version }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/stockBasicInfo';
  import { State, newState } from './model';
  import { adaModalWidth } from '@/utils/hotgo';

  const message = useMessage();
  
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref(newState(null));
  const dialogWidth = computed(() => {
    return adaModalWidth(580);
  });

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    loading.value = true;
    View({ id: state.id })
      .then((res) => {
        formValue.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less" scoped></style>