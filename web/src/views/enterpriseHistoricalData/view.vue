<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="企业级历史行情数据表 (K线数据)详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                证券代码 (如: 000001.SZ, AAPL)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易时间 (精确到分钟或日)
              </template>
              {{ formValue.t }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易日期 (衍生字段)
              </template>
              {{ formValue.date }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易年份
              </template>
              {{ formValue.year }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易月份
              </template>
              {{ formValue.month }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                星期几 (1=周日,7=周六)
              </template>
              {{ formValue.weekday }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                开盘价
              </template>
              {{ formValue.o }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最高价
              </template>
              {{ formValue.h }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                最低价
              </template>
              {{ formValue.l }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                收盘价
              </template>
              {{ formValue.c }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                前收盘价
              </template>
              {{ formValue.pc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                成交量 (股/手)
              </template>
              {{ formValue.v }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                成交额 (元)
              </template>
              {{ formValue.a }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                涨跌额
              </template>
              {{ formValue.change }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                涨跌幅 (%)
              </template>
              {{ formValue.changePct }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                振幅 (%)
              </template>
              {{ formValue.amplitude }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                停牌标志: 0-正常, 1-停牌
              </template>
              {{ formValue.sf }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易状态描述
              </template>
              {{ formValue.tradingStatus }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据周期: 1min, 5min, 15min, 30min, 60min, day, week, month
              </template>
              {{ formValue.period }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否复权: 0-不复权, 1-前复权, 2-后复权
              </template>
              {{ formValue.isAdjusted }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据质量: 0-异常, 1-正常, 2-补全
              </template>
              {{ formValue.dataQuality }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否已验证: 0-未验证, 1-已验证
              </template>
              {{ formValue.isVerified }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据版本
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
  import { View } from '@/api/enterpriseHistoricalData';
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