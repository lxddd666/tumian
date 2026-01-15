<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑资产负债表 #' + formValue.id : '添加资产负债表'"
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
                <n-form-item label="公司代码/股票代码 (例如: 000001.SZ, AAPL)" path="symbol">
                  <n-input
                    placeholder="请输入公司代码/股票代码 (例如: 000001.SZ, AAPL)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="截止日期 (会计期间结束日)" path="jzrq">
                  <DatePicker v-model:formValue="formValue.jzrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="披露日期" path="plrq">
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
                <n-form-item label="报告季度 (1-4, 年报为NULL)" path="reportQuarter">
                  <n-input-number
                    placeholder="请输入报告季度 (1-4, 年报为NULL)"
                    v-model:value="formValue.reportQuarter"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告类型: annual-年报, quarter-季报, interim-中报" path="reportType">
                  <n-input
                    placeholder="请输入报告类型: annual-年报, quarter-季报, interim-中报"
                    v-model:value="formValue.reportType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="货币资金" path="hbzj">
                  <n-input-number
                    placeholder="请输入货币资金"
                    v-model:value="formValue.hbzj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易性金融资产" path="jyxjrzc">
                  <n-input-number
                    placeholder="请输入交易性金融资产"
                    v-model:value="formValue.jyxjrzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应收票据" path="yspj">
                  <n-input-number
                    placeholder="请输入应收票据"
                    v-model:value="formValue.yspj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应收账款" path="yszk">
                  <n-input-number
                    placeholder="请输入应收账款"
                    v-model:value="formValue.yszk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预付款项" path="yfkx">
                  <n-input-number
                    placeholder="请输入预付款项"
                    v-model:value="formValue.yfkx"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应收利息" path="yslx">
                  <n-input-number
                    placeholder="请输入应收利息"
                    v-model:value="formValue.yslx"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应收股利" path="ysgl">
                  <n-input-number
                    placeholder="请输入应收股利"
                    v-model:value="formValue.ysgl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他应收款" path="qtysk">
                  <n-input-number
                    placeholder="请输入其他应收款"
                    v-model:value="formValue.qtysk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="存货" path="ch">
                  <n-input-number
                    placeholder="请输入存货"
                    v-model:value="formValue.ch"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="待摊费用" path="dfy">
                  <n-input-number
                    placeholder="请输入待摊费用"
                    v-model:value="formValue.dfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="一年内到期的非流动资产" path="ynndqdfldzc">
                  <n-input-number
                    placeholder="请输入一年内到期的非流动资产"
                    v-model:value="formValue.ynndqdfldzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他流动资产" path="qtldzc">
                  <n-input-number
                    placeholder="请输入其他流动资产"
                    v-model:value="formValue.qtldzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="流动资产合计" path="ldzchj">
                  <n-input-number
                    placeholder="请输入流动资产合计"
                    v-model:value="formValue.ldzchj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="长期股权投资" path="cqgqtz">
                  <n-input-number
                    placeholder="请输入长期股权投资"
                    v-model:value="formValue.cqgqtz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="长期应收款" path="cqysk">
                  <n-input-number
                    placeholder="请输入长期应收款"
                    v-model:value="formValue.cqysk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="固定资产" path="gdzc">
                  <n-input-number
                    placeholder="请输入固定资产"
                    v-model:value="formValue.gdzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="在建工程" path="zjgc">
                  <n-input-number
                    placeholder="请输入在建工程"
                    v-model:value="formValue.zjgc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="无形资产" path="wxzc">
                  <n-input-number
                    placeholder="请输入无形资产"
                    v-model:value="formValue.wxzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="商誉" path="sy">
                  <n-input-number
                    placeholder="请输入商誉"
                    v-model:value="formValue.sy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="长期待摊费用" path="cqdtfy">
                  <n-input-number
                    placeholder="请输入长期待摊费用"
                    v-model:value="formValue.cqdtfy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="递延所得税资产" path="dysdszc">
                  <n-input-number
                    placeholder="请输入递延所得税资产"
                    v-model:value="formValue.dysdszc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他非流动资产" path="qtfldzc">
                  <n-input-number
                    placeholder="请输入其他非流动资产"
                    v-model:value="formValue.qtfldzc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="非流动资产合计" path="fldzchj">
                  <n-input-number
                    placeholder="请输入非流动资产合计"
                    v-model:value="formValue.fldzchj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="资产总计" path="zczj">
                  <n-input-number
                    placeholder="请输入资产总计"
                    v-model:value="formValue.zczj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="短期借款" path="dqjk">
                  <n-input-number
                    placeholder="请输入短期借款"
                    v-model:value="formValue.dqjk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易性金融负债" path="jyxjrfz">
                  <n-input-number
                    placeholder="请输入交易性金融负债"
                    v-model:value="formValue.jyxjrfz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应付票据" path="yfpj">
                  <n-input-number
                    placeholder="请输入应付票据"
                    v-model:value="formValue.yfpj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应付账款" path="yfzk">
                  <n-input-number
                    placeholder="请输入应付账款"
                    v-model:value="formValue.yfzk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预收账款" path="ysk">
                  <n-input-number
                    placeholder="请输入预收账款"
                    v-model:value="formValue.ysk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应付职工薪酬" path="yfgzxc">
                  <n-input-number
                    placeholder="请输入应付职工薪酬"
                    v-model:value="formValue.yfgzxc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应交税费" path="yjsf">
                  <n-input-number
                    placeholder="请输入应交税费"
                    v-model:value="formValue.yjsf"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应付利息" path="yflx">
                  <n-input-number
                    placeholder="请输入应付利息"
                    v-model:value="formValue.yflx"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应付股利" path="yfgl">
                  <n-input-number
                    placeholder="请输入应付股利"
                    v-model:value="formValue.yfgl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他应付款" path="qtfzk">
                  <n-input-number
                    placeholder="请输入其他应付款"
                    v-model:value="formValue.qtfzk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="一年内到期的非流动负债" path="ynndqdfldfz">
                  <n-input-number
                    placeholder="请输入一年内到期的非流动负债"
                    v-model:value="formValue.ynndqdfldfz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他流动负债" path="qtldfz">
                  <n-input-number
                    placeholder="请输入其他流动负债"
                    v-model:value="formValue.qtldfz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="流动负债合计" path="ldfzhj">
                  <n-input-number
                    placeholder="请输入流动负债合计"
                    v-model:value="formValue.ldfzhj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="长期借款" path="cqjk">
                  <n-input-number
                    placeholder="请输入长期借款"
                    v-model:value="formValue.cqjk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="应付债券" path="yfzq">
                  <n-input-number
                    placeholder="请输入应付债券"
                    v-model:value="formValue.yfzq"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="长期应付款" path="cqyfk">
                  <n-input-number
                    placeholder="请输入长期应付款"
                    v-model:value="formValue.cqyfk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="递延所得税负债" path="dysdsfz">
                  <n-input-number
                    placeholder="请输入递延所得税负债"
                    v-model:value="formValue.dysdsfz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他非流动负债" path="qtfldfz">
                  <n-input-number
                    placeholder="请输入其他非流动负债"
                    v-model:value="formValue.qtfldfz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="非流动负债合计" path="fldfzhj">
                  <n-input-number
                    placeholder="请输入非流动负债合计"
                    v-model:value="formValue.fldfzhj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="负债合计" path="fzhj">
                  <n-input-number
                    placeholder="请输入负债合计"
                    v-model:value="formValue.fzhj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="实收资本(或股本)" path="sszb">
                  <n-input-number
                    placeholder="请输入实收资本(或股本)"
                    v-model:value="formValue.sszb"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="资本公积" path="zbgj">
                  <n-input-number
                    placeholder="请输入资本公积"
                    v-model:value="formValue.zbgj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="盈余公积" path="ylgj">
                  <n-input-number
                    placeholder="请输入盈余公积"
                    v-model:value="formValue.ylgj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="未分配利润" path="wfplr">
                  <n-input-number
                    placeholder="请输入未分配利润"
                    v-model:value="formValue.wfplr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="归属于母公司股东权益合计" path="gsmgdqsyhj">
                  <n-input-number
                    placeholder="请输入归属于母公司股东权益合计"
                    v-model:value="formValue.gsmgdqsyhj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="少数股东权益" path="ssgdqy">
                  <n-input-number
                    placeholder="请输入少数股东权益"
                    v-model:value="formValue.ssgdqy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="所有者权益合计" path="syzqyhj">
                  <n-input-number
                    placeholder="请输入所有者权益合计"
                    v-model:value="formValue.syzqyhj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="负债和股东权益总计" path="fzhgdqyzj">
                  <n-input-number
                    placeholder="请输入负债和股东权益总计"
                    v-model:value="formValue.fzhgdqyzj"
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
                <n-form-item label="货币单位 (CNY, USD等)" path="currency">
                  <n-input
                    placeholder="请输入货币单位 (CNY, USD等)"
                    v-model:value="formValue.currency"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="单位: yuan-元, wan-万元, qianwan-千万元" path="unit">
                  <n-input
                    placeholder="请输入单位: yuan-元, wan-万元, qianwan-千万元"
                    v-model:value="formValue.unit"
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
                <n-form-item label="数据版本" path="version">
                  <n-input-number
                    placeholder="请输入数据版本"
                    v-model:value="formValue.version"
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
  import { Edit, View } from '@/api/balanceSheet';
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