<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="基金持股明细表 (来源于基金定期报告)详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                截止日期 (报告期，如2025-12-31)
              </template>
              {{ formValue.jzrq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易时间 (衍生自jzrq，兼容时间序列查询)
              </template>
              {{ formValue.t }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                基金名称
              </template>
              {{ formValue.jjmc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                基金代码
              </template>
              {{ formValue.jjdm }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股票代码 (如: 000001.SZ)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                持仓数量(股)
              </template>
              {{ formValue.ccsl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                占流通股比例(%)
              </template>
              {{ formValue.ltbl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                持股市值（元）
              </template>
              {{ formValue.cgsz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                占净值比例（%）
              </template>
              {{ formValue.jzbl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                估算持仓成本 (元/股)
              </template>
              {{ formValue.avgCost }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源 (如: 基金季报)
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                报告类型: 季报, 中报, 年报
              </template>
              {{ formValue.reportType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否为该基金对该股票的最新持仓: 0-否, 1-是
              </template>
              {{ formValue.isLatest }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/fundStockHolding';
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