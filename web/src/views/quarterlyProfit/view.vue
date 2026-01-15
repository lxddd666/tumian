<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="季度利润数据表 (近一年各季度)详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                股票代码 (如: 000001.SZ)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                截止日期 (报告期截止日，如2025-03-31)
              </template>
              {{ formValue.date }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                报告年度
              </template>
              {{ formValue.reportYear }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                报告季度 (1-4)
              </template>
              {{ formValue.reportQuarter }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                会计期间 (如2025Q1)
              </template>
              {{ formValue.fiscalPeriod }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业收入（万元）
              </template>
              {{ formValue.income }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业支出（万元）
              </template>
              {{ formValue.expend }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业利润（万元）
              </template>
              {{ formValue.profit }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                利润总额（万元）
              </template>
              {{ formValue.totalp }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利润（万元）
              </template>
              {{ formValue.reprofit }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                基本每股收益(元/股)
              </template>
              {{ formValue.basege }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                稀释每股收益(元/股)
              </template>
              {{ formValue.ettege }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他综合收益（万元）
              </template>
              {{ formValue.otherp }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                综合收益总额（万元）
              </template>
              {{ formValue.totalcp }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                毛利率(%)
              </template>
              {{ formValue.grossProfitMargin }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利率(%)
              </template>
              {{ formValue.netProfitMargin }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业利润率(%)
              </template>
              {{ formValue.operatingProfitRatio }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                报告类型: 一季报, 中报, 三季报, 年报
              </template>
              {{ formValue.reportType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源 (如: 交易所财报)
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                货币单位
              </template>
              {{ formValue.currency }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否审计: 0-未审计, 1-已审计
              </template>
              {{ formValue.isAudited }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/quarterlyProfit';
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