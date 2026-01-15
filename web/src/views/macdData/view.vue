<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="MACD指标数据表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                股票或标的代码 (例如: AAPL, 000001.SH)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)
              </template>
              {{ formValue.t }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据间隔: minute-短分时, day-日线
              </template>
              {{ formValue.intervalType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                DIFF值
              </template>
              {{ formValue.diff }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                DEA值
              </template>
              {{ formValue.dea }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                MACD值
              </template>
              {{ formValue.macd }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                EMA(12)值
              </template>
              {{ formValue.ema12 }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                EMA(26)值
              </template>
              {{ formValue.ema26 }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/macdData';
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