<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content title="利润表 (Income Statement)详情" closable>
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
                截止日期 (报告期截止日，如2025-12-31)
              </template>
              {{ formValue.jzrq }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                披露日期 (财报实际发布日期)
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
                会计期间
              </template>
              {{ formValue.fiscalPeriod }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业收入
              </template>
              {{ formValue.yysr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                已赚保费
              </template>
              {{ formValue.yzbf }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                房地产销售收入
              </template>
              {{ formValue.fdczssr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他业务收入
              </template>
              {{ formValue.qtywsr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业总收入
              </template>
              {{ formValue.yyzsr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                利息收入
              </template>
              {{ formValue.lxsr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                手续费及佣金收入
              </template>
              {{ formValue.sxfjyjsr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                补贴收入
              </template>
              {{ formValue.btsr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业外收入
              </template>
              {{ formValue.ywsr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他收益
              </template>
              {{ formValue.qtsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业成本
              </template>
              {{ formValue.yycb }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                房地产销售成本
              </template>
              {{ formValue.fdczscb }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他业务成本
              </template>
              {{ formValue.qtywcb }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业总成本
              </template>
              {{ formValue.yyzcb }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业税金及附加
              </template>
              {{ formValue.yysjjfj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                销售费用
              </template>
              {{ formValue.xsfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                管理费用
              </template>
              {{ formValue.glfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                研发费用
              </template>
              {{ formValue.yffy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                财务费用
              </template>
              {{ formValue.cwfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                手续费及佣金支出
              </template>
              {{ formValue.sxfjyjzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                利息支出
              </template>
              {{ formValue.lxzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                退保金
              </template>
              {{ formValue.tbj }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                赔付支出净额
              </template>
              {{ formValue.pczjje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                提取保险合同准备金净额
              </template>
              {{ formValue.tqbxhtzbjje }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                保单红利支出
              </template>
              {{ formValue.bdhlzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                分保费用
              </template>
              {{ formValue.fbfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                资产减值损失
              </template>
              {{ formValue.zcjzss }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业外支出
              </template>
              {{ formValue.ywzc }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                其他业务利润
              </template>
              {{ formValue.qtywlr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业利润
              </template>
              {{ formValue.yylr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                利润总额
              </template>
              {{ formValue.lrze }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利润
              </template>
              {{ formValue.jlr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利润(扣除非经常性损益后)
              </template>
              {{ formValue.jlrhfcjcx }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                归属于母公司所有者的净利润
              </template>
              {{ formValue.gsmgsyzzdjlr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                被合并方在合并前实现净利润
              </template>
              {{ formValue.bhbfzhbqsljlr }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                投资收益
              </template>
              {{ formValue.tzsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                联营企业和合营企业的投资收益
              </template>
              {{ formValue.lyqyhhhqydtzsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                公允价值变动收益
              </template>
              {{ formValue.gyjzbdsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                期货损益
              </template>
              {{ formValue.qhsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                托管收益
              </template>
              {{ formValue.tgsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                汇兑收益
              </template>
              {{ formValue.hdsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                非流动资产处置收益
              </template>
              {{ formValue.fldzcczsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                所得税费用
              </template>
              {{ formValue.sdsfy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                少数股东损益
              </template>
              {{ formValue.ssgdsy }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                未确认投资损失
              </template>
              {{ formValue.wqrtzss }}
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
                综合收益总额
              </template>
              {{ formValue.zhsyz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                归属于少数股东的综合收益总额
              </template>
              {{ formValue.gsssgdzhsyz }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                毛利率(%)
              </template>
              {{ formValue.grossMargin }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                营业利润率(%)
              </template>
              {{ formValue.operatingMargin }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                净利率(%)
              </template>
              {{ formValue.netMargin }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                实际税率(%)
              </template>
              {{ formValue.effectiveTaxRate }}
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
                单位: yuan-元, wan-万元
              </template>
              {{ formValue.unit }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                会计准则 (如: CAS, IFRS)
              </template>
              {{ formValue.accountingStandard }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否审计: 0-未审计, 1-已审计
              </template>
              {{ formValue.isAudited }}
            </n-descriptions-item>
            <n-descriptions-item>
              <template #label>
                是否合并报表: 1-合并, 0-母公司
              </template>
              {{ formValue.isConsolidated }}
            </n-descriptions-item>
          </n-descriptions>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>
<script lang="ts" setup>
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/incomeStatement';
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