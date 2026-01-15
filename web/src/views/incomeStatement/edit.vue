<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑利润表 (Income Statement) #' + formValue.id : '添加利润表 (Income Statement)'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid
              cols="1 s:1 m:1 l:1 xl:1 2xl:1"
              responsive="screen"
            >
              <n-gi span="1">
                <n-form-item label="股票代码 (如: 000001.SZ)" path="symbol">
                  <n-input
                    placeholder="请输入股票代码 (如: 000001.SZ)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="截止日期 (报告期截止日，如2025-12-31)" path="jzrq">
                  <DatePicker v-model:formValue="formValue.jzrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="披露日期 (财报实际发布日期)" path="plrq">
                  <DatePicker v-model:formValue="formValue.plrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告年度" path="reportYear">
                  <n-input-number
                    placeholder="请输入报告年度"
                    v-model:value="formValue.reportYear"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告季度 (1-4)" path="reportQuarter">
                  <n-input-number
                    placeholder="请输入报告季度 (1-4)"
                    v-model:value="formValue.reportQuarter"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告类型: annual-年报, quarter-季报" path="reportType">
                  <n-input
                    placeholder="请输入报告类型: annual-年报, quarter-季报"
                    v-model:value="formValue.reportType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="会计期间" path="fiscalPeriod">
                  <n-input
                    placeholder="请输入会计期间"
                    v-model:value="formValue.fiscalPeriod"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业收入" path="yysr">
                  <n-input-number
                    placeholder="请输入营业收入"
                    v-model:value="formValue.yysr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="已赚保费" path="yzbf">
                  <n-input-number
                    placeholder="请输入已赚保费"
                    v-model:value="formValue.yzbf"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="房地产销售收入" path="fdczssr">
                  <n-input-number
                    placeholder="请输入房地产销售收入"
                    v-model:value="formValue.fdczssr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他业务收入" path="qtywsr">
                  <n-input-number
                    placeholder="请输入其他业务收入"
                    v-model:value="formValue.qtywsr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业总收入" path="yyzsr">
                  <n-input-number
                    placeholder="请输入营业总收入"
                    v-model:value="formValue.yyzsr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="利息收入" path="lxsr">
                  <n-input-number
                    placeholder="请输入利息收入"
                    v-model:value="formValue.lxsr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手续费及佣金收入" path="sxfjyjsr">
                  <n-input-number
                    placeholder="请输入手续费及佣金收入"
                    v-model:value="formValue.sxfjyjsr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="补贴收入" path="btsr">
                  <n-input-number
                    placeholder="请输入补贴收入"
                    v-model:value="formValue.btsr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业外收入" path="ywsr">
                  <n-input-number
                    placeholder="请输入营业外收入"
                    v-model:value="formValue.ywsr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他收益" path="qtsy">
                  <n-input-number
                    placeholder="请输入其他收益"
                    v-model:value="formValue.qtsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业成本" path="yycb">
                  <n-input-number
                    placeholder="请输入营业成本"
                    v-model:value="formValue.yycb"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="房地产销售成本" path="fdczscb">
                  <n-input-number
                    placeholder="请输入房地产销售成本"
                    v-model:value="formValue.fdczscb"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他业务成本" path="qtywcb">
                  <n-input-number
                    placeholder="请输入其他业务成本"
                    v-model:value="formValue.qtywcb"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业总成本" path="yyzcb">
                  <n-input-number
                    placeholder="请输入营业总成本"
                    v-model:value="formValue.yyzcb"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业税金及附加" path="yysjjfj">
                  <n-input-number
                    placeholder="请输入营业税金及附加"
                    v-model:value="formValue.yysjjfj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="销售费用" path="xsfy">
                  <n-input-number
                    placeholder="请输入销售费用"
                    v-model:value="formValue.xsfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="管理费用" path="glfy">
                  <n-input-number
                    placeholder="请输入管理费用"
                    v-model:value="formValue.glfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="研发费用" path="yffy">
                  <n-input-number
                    placeholder="请输入研发费用"
                    v-model:value="formValue.yffy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="财务费用" path="cwfy">
                  <n-input-number
                    placeholder="请输入财务费用"
                    v-model:value="formValue.cwfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="手续费及佣金支出" path="sxfjyjzc">
                  <n-input-number
                    placeholder="请输入手续费及佣金支出"
                    v-model:value="formValue.sxfjyjzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="利息支出" path="lxzc">
                  <n-input-number
                    placeholder="请输入利息支出"
                    v-model:value="formValue.lxzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退保金" path="tbj">
                  <n-input-number
                    placeholder="请输入退保金"
                    v-model:value="formValue.tbj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="赔付支出净额" path="pczjje">
                  <n-input-number
                    placeholder="请输入赔付支出净额"
                    v-model:value="formValue.pczjje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="提取保险合同准备金净额" path="tqbxhtzbjje">
                  <n-input-number
                    placeholder="请输入提取保险合同准备金净额"
                    v-model:value="formValue.tqbxhtzbjje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="保单红利支出" path="bdhlzc">
                  <n-input-number
                    placeholder="请输入保单红利支出"
                    v-model:value="formValue.bdhlzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="分保费用" path="fbfy">
                  <n-input-number
                    placeholder="请输入分保费用"
                    v-model:value="formValue.fbfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="资产减值损失" path="zcjzss">
                  <n-input-number
                    placeholder="请输入资产减值损失"
                    v-model:value="formValue.zcjzss"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业外支出" path="ywzc">
                  <n-input-number
                    placeholder="请输入营业外支出"
                    v-model:value="formValue.ywzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他业务利润" path="qtywlr">
                  <n-input-number
                    placeholder="请输入其他业务利润"
                    v-model:value="formValue.qtywlr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业利润" path="yylr">
                  <n-input-number
                    placeholder="请输入营业利润"
                    v-model:value="formValue.yylr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="利润总额" path="lrze">
                  <n-input-number
                    placeholder="请输入利润总额"
                    v-model:value="formValue.lrze"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利润" path="jlr">
                  <n-input-number
                    placeholder="请输入净利润"
                    v-model:value="formValue.jlr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利润(扣除非经常性损益后)" path="jlrhfcjcx">
                  <n-input-number
                    placeholder="请输入净利润(扣除非经常性损益后)"
                    v-model:value="formValue.jlrhfcjcx"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="归属于母公司所有者的净利润" path="gsmgsyzzdjlr">
                  <n-input-number
                    placeholder="请输入归属于母公司所有者的净利润"
                    v-model:value="formValue.gsmgsyzzdjlr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被合并方在合并前实现净利润" path="bhbfzhbqsljlr">
                  <n-input-number
                    placeholder="请输入被合并方在合并前实现净利润"
                    v-model:value="formValue.bhbfzhbqsljlr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="投资收益" path="tzsy">
                  <n-input-number
                    placeholder="请输入投资收益"
                    v-model:value="formValue.tzsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联营企业和合营企业的投资收益" path="lyqyhhhqydtzsy">
                  <n-input-number
                    placeholder="请输入联营企业和合营企业的投资收益"
                    v-model:value="formValue.lyqyhhhqydtzsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="公允价值变动收益" path="gyjzbdsy">
                  <n-input-number
                    placeholder="请输入公允价值变动收益"
                    v-model:value="formValue.gyjzbdsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="期货损益" path="qhsy">
                  <n-input-number
                    placeholder="请输入期货损益"
                    v-model:value="formValue.qhsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="托管收益" path="tgsy">
                  <n-input-number
                    placeholder="请输入托管收益"
                    v-model:value="formValue.tgsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="汇兑收益" path="hdsy">
                  <n-input-number
                    placeholder="请输入汇兑收益"
                    v-model:value="formValue.hdsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="非流动资产处置收益" path="fldzcczsy">
                  <n-input-number
                    placeholder="请输入非流动资产处置收益"
                    v-model:value="formValue.fldzcczsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="所得税费用" path="sdsfy">
                  <n-input-number
                    placeholder="请输入所得税费用"
                    v-model:value="formValue.sdsfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="少数股东损益" path="ssgdsy">
                  <n-input-number
                    placeholder="请输入少数股东损益"
                    v-model:value="formValue.ssgdsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="未确认投资损失" path="wqrtzss">
                  <n-input-number
                    placeholder="请输入未确认投资损失"
                    v-model:value="formValue.wqrtzss"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="基本每股收益" path="jbmgsy">
                  <n-input-number
                    placeholder="请输入基本每股收益"
                    v-model:value="formValue.jbmgsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="稀释每股收益" path="xsmgsy">
                  <n-input-number
                    placeholder="请输入稀释每股收益"
                    v-model:value="formValue.xsmgsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="综合收益总额" path="zhsyz">
                  <n-input-number
                    placeholder="请输入综合收益总额"
                    v-model:value="formValue.zhsyz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="归属于少数股东的综合收益总额" path="gsssgdzhsyz">
                  <n-input-number
                    placeholder="请输入归属于少数股东的综合收益总额"
                    v-model:value="formValue.gsssgdzhsyz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="毛利率(%)" path="grossMargin">
                  <n-input-number
                    placeholder="请输入毛利率(%)"
                    v-model:value="formValue.grossMargin"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业利润率(%)" path="operatingMargin">
                  <n-input-number
                    placeholder="请输入营业利润率(%)"
                    v-model:value="formValue.operatingMargin"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利率(%)" path="netMargin">
                  <n-input-number
                    placeholder="请输入净利率(%)"
                    v-model:value="formValue.netMargin"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="实际税率(%)" path="effectiveTaxRate">
                  <n-input-number
                    placeholder="请输入实际税率(%)"
                    v-model:value="formValue.effectiveTaxRate"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据来源" path="dataSource">
                  <n-input
                    placeholder="请输入数据来源"
                    v-model:value="formValue.dataSource"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="货币单位" path="currency">
                  <n-input
                    placeholder="请输入货币单位"
                    v-model:value="formValue.currency"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="单位: yuan-元, wan-万元" path="unit">
                  <n-input
                    placeholder="请输入单位: yuan-元, wan-万元"
                    v-model:value="formValue.unit"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="会计准则 (如: CAS, IFRS)" path="accountingStandard">
                  <n-input
                    placeholder="请输入会计准则 (如: CAS, IFRS)"
                    v-model:value="formValue.accountingStandard"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否审计: 0-未审计, 1-已审计" path="isAudited">
                  <n-input-number
                    placeholder="请输入是否审计: 0-未审计, 1-已审计"
                    v-model:value="formValue.isAudited"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否合并报表: 1-合并, 0-母公司" path="isConsolidated">
                  <n-input-number
                    placeholder="请输入是否合并报表: 1-合并, 0-母公司"
                    v-model:value="formValue.isConsolidated"
                    />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" :disabled="!isFormValid" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>
<script lang="ts" setup>
  import { Edit, View } from '@/api/incomeStatement';
  import { State, newState, rules } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = ref<State>(newState(null));
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });
  const isFormValid = ref(true);

  // 提交表单
  function confirmForm(e) {
    e.preventDefault();
    formRef.value.validate((errors) => {
      if (!errors) {
        formBtnLoading.value = true;
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            closeForm();
            emit('reloadTable');
          })
          .finally(() => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('请填写完整信息');
      }
    });
  }

  // 关闭表单
  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  // 打开模态框
  function openModal(state: State) {
    showModal.value = true;
    
    // 新增
    if (!state || state.id < 1) {
      formValue.value = newState(state);
      
      return;
    }

    // 编辑
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

<style lang="less"></style>