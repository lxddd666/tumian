<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="ai 选股判断详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                股票或标的代码 (例如: AAPL, 000001.SZ)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                时间
              </template>
              {{ formValue.t }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                ai id
              </template>
              {{ formValue.aiId }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                指标判断
              </template>
              <span v-html="formValue.indicatorsJudgment"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                指标判断 1是2否
              </template>
              {{ formValue.indicatorsFlag }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                财务判断
              </template>
              <span v-html="formValue.financialJudgment"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                财务判断 1是2否
              </template>
              {{ formValue.financialFlag }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                综合判断
              </template>
              <span v-html="formValue.comprehensiveJudgment"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                综合判断 1是2否
              </template>
              {{ formValue.comprehensiveFlag }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/stockAiJudgment';
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