<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="财务指标分析表详情" closable>
        <n-spin :show="loading" description="请稍候...">
          <n-descriptions label-placement="left" class="py-2" :column="1">
            <n-descriptions-item>
              <template #label>
                公司代码/股票代码 (例如: 000001.SZ, AAPL)
              </template>
              {{ formValue.symbol }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                截止日期 (报告期结束日)
              </template>
              {{ formValue.jzrq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                披露日期
              </template>
              {{ formValue.plrq }}
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
                报告类型: annual-年报, quarter-季报
              </template>
              {{ formValue.reportType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                会计期间 (衍生字段，如2023Q1)
              </template>
              {{ formValue.fiscalPeriod }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                每股资本公积金
              </template>
              {{ formValue.mgzbgjj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                每股经营活动现金流量
              </template>
              {{ formValue.mgjyhdxjl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                每股净资产
              </template>
              {{ formValue.mgjzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                基本每股收益
              </template>
              {{ formValue.jbmgsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                稀释每股收益
              </template>
              {{ formValue.xsmgsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                每股未分配利润
              </template>
              {{ formValue.mgwfplr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                扣非每股收益
              </template>
              {{ formValue.kfmgsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净资产收益率(%)
              </template>
              {{ formValue.jzcsyl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                加权净资产收益率(%)
              </template>
              {{ formValue.jqjzcsyl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                摊薄净资产收益率(%)
              </template>
              {{ formValue.tbjzcsyl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                摊薄总资产收益率(%)
              </template>
              {{ formValue.tbzzcsyl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                销售毛利率(%)
              </template>
              {{ formValue.xsmlv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                毛利率(%)
              </template>
              {{ formValue.mlv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利率(%)
              </template>
              {{ formValue.jlv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                实际税率(%)
              </template>
              {{ formValue.sjslv }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                主营收入同比增长(%)
              </template>
              {{ formValue.zyyrsrzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利润同比增长(%)
              </template>
              {{ formValue.jlrzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                归属于母公司所有者的净利润同比增长(%)
              </template>
              {{ formValue.gsmgsyzzdjlrzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                扣非净利润同比增长(%)
              </template>
              {{ formValue.kfjlrzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业总收入滚动环比增长(%)
              </template>
              {{ formValue.yyzsrgdhbzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                归属净利润滚动环比增长(%)
              </template>
              {{ formValue.sljlrjqhbzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                扣非净利润滚动环比增长(%)
              </template>
              {{ formValue.kfjlrgdhbzz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预收款/营业收入
              </template>
              {{ formValue.yskyysr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                销售现金流/营业收入
              </template>
              {{ formValue.xsxjlyysr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                资产负债比率(%)
              </template>
              {{ formValue.zcfzl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                存货周转率(次)
              </template>
              {{ formValue.chzzl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源
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
                单位: yuan-元
              </template>
              {{ formValue.unit }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否为计算指标: 0-原始数据, 1-计算得出
              </template>
              {{ formValue.isCalculated }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                计算版本
              </template>
              {{ formValue.calcVersion }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/financialIndicators';
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