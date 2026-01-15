<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="资产负债表详情" closable>
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
                截止日期 (会计期间结束日)
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
                报告季度 (1-4, 年报为NULL)
              </template>
              {{ formValue.reportQuarter }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                报告类型: annual-年报, quarter-季报, interim-中报
              </template>
              {{ formValue.reportType }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                货币资金
              </template>
              {{ formValue.hbzj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易性金融资产
              </template>
              {{ formValue.jyxjrzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应收票据
              </template>
              {{ formValue.yspj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应收账款
              </template>
              {{ formValue.yszk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预付款项
              </template>
              {{ formValue.yfkx }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应收利息
              </template>
              {{ formValue.yslx }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应收股利
              </template>
              {{ formValue.ysgl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他应收款
              </template>
              {{ formValue.qtysk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                存货
              </template>
              {{ formValue.ch }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                待摊费用
              </template>
              {{ formValue.dfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                一年内到期的非流动资产
              </template>
              {{ formValue.ynndqdfldzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他流动资产
              </template>
              {{ formValue.qtldzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                流动资产合计
              </template>
              {{ formValue.ldzchj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                长期股权投资
              </template>
              {{ formValue.cqgqtz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                长期应收款
              </template>
              {{ formValue.cqysk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                固定资产
              </template>
              {{ formValue.gdzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                在建工程
              </template>
              {{ formValue.zjgc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                无形资产
              </template>
              {{ formValue.wxzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                商誉
              </template>
              {{ formValue.sy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                长期待摊费用
              </template>
              {{ formValue.cqdtfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                递延所得税资产
              </template>
              {{ formValue.dysdszc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他非流动资产
              </template>
              {{ formValue.qtfldzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                非流动资产合计
              </template>
              {{ formValue.fldzchj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                资产总计
              </template>
              {{ formValue.zczj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                短期借款
              </template>
              {{ formValue.dqjk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                交易性金融负债
              </template>
              {{ formValue.jyxjrfz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应付票据
              </template>
              {{ formValue.yfpj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应付账款
              </template>
              {{ formValue.yfzk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                预收账款
              </template>
              {{ formValue.ysk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应付职工薪酬
              </template>
              {{ formValue.yfgzxc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应交税费
              </template>
              {{ formValue.yjsf }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应付利息
              </template>
              {{ formValue.yflx }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应付股利
              </template>
              {{ formValue.yfgl }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他应付款
              </template>
              {{ formValue.qtfzk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                一年内到期的非流动负债
              </template>
              {{ formValue.ynndqdfldfz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他流动负债
              </template>
              {{ formValue.qtldfz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                流动负债合计
              </template>
              {{ formValue.ldfzhj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                长期借款
              </template>
              {{ formValue.cqjk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                应付债券
              </template>
              {{ formValue.yfzq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                长期应付款
              </template>
              {{ formValue.cqyfk }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                递延所得税负债
              </template>
              {{ formValue.dysdsfz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他非流动负债
              </template>
              {{ formValue.qtfldfz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                非流动负债合计
              </template>
              {{ formValue.fldfzhj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                负债合计
              </template>
              {{ formValue.fzhj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                实收资本(或股本)
              </template>
              {{ formValue.sszb }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                资本公积
              </template>
              {{ formValue.zbgj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                盈余公积
              </template>
              {{ formValue.ylgj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                未分配利润
              </template>
              {{ formValue.wfplr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                归属于母公司股东权益合计
              </template>
              {{ formValue.gsmgdqsyhj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                少数股东权益
              </template>
              {{ formValue.ssgdqy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                所有者权益合计
              </template>
              {{ formValue.syzqyhj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                负债和股东权益总计
              </template>
              {{ formValue.fzhgdqyzj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                数据来源
              </template>
              {{ formValue.dataSource }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                货币单位 (CNY, USD等)
              </template>
              {{ formValue.currency }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                单位: yuan-元, wan-万元, qianwan-千万元
              </template>
              {{ formValue.unit }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否审计: 0-未审计, 1-已审计
              </template>
              {{ formValue.isAudited }}
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
  import { View } from '@/api/balanceSheet';
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