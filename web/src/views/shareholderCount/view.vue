<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="公司股东户数统计表 (按报告期统计)详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                公司代码/股票代码 (例如: 000001.SZ, 600000.SS)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                截止日期 (报告期结束日, 如2023-09-30)
              </template>
              {{ formValue.jzrq }}
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
                报告类型: annual-年报, half_year-中报, quarter-季报
              </template>
              {{ formValue.reportType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股东总数 (户)
              </template>
              {{ formValue.gdzs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                A股东户数 (户)
              </template>
              {{ formValue.agdhs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                B股东户数 (户)
              </template>
              {{ formValue.bgdhs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                H股东户数 (户)
              </template>
              {{ formValue.hgdhs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                已流通股东户数 (户)
              </template>
              {{ formValue.yltgdhs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                未流通股东户数 (户)
              </template>
              {{ formValue.wltgdhs }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                A股股东占比(%)
              </template>
              {{ formValue.agRatio }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                已流通股东占比(%)
              </template>
              {{ formValue.yltRatio }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否为该报告期最新数据
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
  import { View } from '@/api/shareholderCount';
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