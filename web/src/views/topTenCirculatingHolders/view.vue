<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
<n-drawer-content title="公司十大流通股东表 (数据来源于定期报告)[citation:4]详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                公司代码/股票代码 (例如: 000001.SZ)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                截止日期 (报告期结束日, 如2023-09-30)[citation:4]
              </template>
              {{ formValue.jzrq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                公告日期 (信息发布日期)
              </template>
              {{ formValue.ggrq }}
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
                报告类型: annual-年报, half_year-中报, quarter-季报[citation:4]
              </template>
              {{ formValue.reportType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股东名称
              </template>
              {{ formValue.gdmc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股东类型 (如: 基金、社保、个人等)
              </template>
              {{ formValue.gdlx }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                股份性质 (如: 流通A股、限售A股等)
              </template>
              {{ formValue.gfxz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                持股数量 (股)
              </template>
              {{ formValue.cgsl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                持股比例 (%)
              </template>
              {{ formValue.cgbl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                持股排名 (1-10)
              </template>
              {{ formValue.cgpm }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                变动原因
              </template>
              <span v-html="formValue.bdyy"></span>
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                变动类型: increase-增持, decrease-减持, new-新进, unchanged-不变, other-其他
              </template>
              {{ formValue.bdType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源 (如: 交易所公告)[citation:4]
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否为该报告期最新数据: 0-历史快照, 1-最新
              </template>
              {{ formValue.isLatest }}
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
  import { View } from '@/api/topTenCirculatingHolders';
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