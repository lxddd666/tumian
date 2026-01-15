<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑财务指标分析表 #' + formValue.id : '添加财务指标分析表'"
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
                <n-form-item label="截止日期 (报告期结束日)" path="jzrq">
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
                <n-form-item label="会计期间 (衍生字段，如2023Q1)" path="fiscalPeriod">
                  <n-input
                    placeholder="请输入会计期间 (衍生字段，如2023Q1)"
                    v-model:value="formValue.fiscalPeriod"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="每股资本公积金" path="mgzbgjj">
                  <n-input-number
                    placeholder="请输入每股资本公积金"
                    v-model:value="formValue.mgzbgjj"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="每股经营活动现金流量" path="mgjyhdxjl">
                  <n-input-number
                    placeholder="请输入每股经营活动现金流量"
                    v-model:value="formValue.mgjyhdxjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="每股净资产" path="mgjzc">
                  <n-input-number
                    placeholder="请输入每股净资产"
                    v-model:value="formValue.mgjzc"
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
                <n-form-item label="每股未分配利润" path="mgwfplr">
                  <n-input-number
                    placeholder="请输入每股未分配利润"
                    v-model:value="formValue.mgwfplr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="扣非每股收益" path="kfmgsy">
                  <n-input-number
                    placeholder="请输入扣非每股收益"
                    v-model:value="formValue.kfmgsy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净资产收益率(%)" path="jzcsyl">
                  <n-input-number
                    placeholder="请输入净资产收益率(%)"
                    v-model:value="formValue.jzcsyl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="加权净资产收益率(%)" path="jqjzcsyl">
                  <n-input-number
                    placeholder="请输入加权净资产收益率(%)"
                    v-model:value="formValue.jqjzcsyl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="摊薄净资产收益率(%)" path="tbjzcsyl">
                  <n-input-number
                    placeholder="请输入摊薄净资产收益率(%)"
                    v-model:value="formValue.tbjzcsyl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="摊薄总资产收益率(%)" path="tbzzcsyl">
                  <n-input-number
                    placeholder="请输入摊薄总资产收益率(%)"
                    v-model:value="formValue.tbzzcsyl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="销售毛利率(%)" path="xsmlv">
                  <n-input-number
                    placeholder="请输入销售毛利率(%)"
                    v-model:value="formValue.xsmlv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="毛利率(%)" path="mlv">
                  <n-input-number
                    placeholder="请输入毛利率(%)"
                    v-model:value="formValue.mlv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利率(%)" path="jlv">
                  <n-input-number
                    placeholder="请输入净利率(%)"
                    v-model:value="formValue.jlv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="实际税率(%)" path="sjslv">
                  <n-input-number
                    placeholder="请输入实际税率(%)"
                    v-model:value="formValue.sjslv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主营收入同比增长(%)" path="zyyrsrzz">
                  <n-input-number
                    placeholder="请输入主营收入同比增长(%)"
                    v-model:value="formValue.zyyrsrzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利润同比增长(%)" path="jlrzz">
                  <n-input-number
                    placeholder="请输入净利润同比增长(%)"
                    v-model:value="formValue.jlrzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="归属于母公司所有者的净利润同比增长(%)" path="gsmgsyzzdjlrzz">
                  <n-input-number
                    placeholder="请输入归属于母公司所有者的净利润同比增长(%)"
                    v-model:value="formValue.gsmgsyzzdjlrzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="扣非净利润同比增长(%)" path="kfjlrzz">
                  <n-input-number
                    placeholder="请输入扣非净利润同比增长(%)"
                    v-model:value="formValue.kfjlrzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业总收入滚动环比增长(%)" path="yyzsrgdhbzz">
                  <n-input-number
                    placeholder="请输入营业总收入滚动环比增长(%)"
                    v-model:value="formValue.yyzsrgdhbzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="归属净利润滚动环比增长(%)" path="sljlrjqhbzz">
                  <n-input-number
                    placeholder="请输入归属净利润滚动环比增长(%)"
                    v-model:value="formValue.sljlrjqhbzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="扣非净利润滚动环比增长(%)" path="kfjlrgdhbzz">
                  <n-input-number
                    placeholder="请输入扣非净利润滚动环比增长(%)"
                    v-model:value="formValue.kfjlrgdhbzz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预收款/营业收入" path="yskyysr">
                  <n-input-number
                    placeholder="请输入预收款/营业收入"
                    v-model:value="formValue.yskyysr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="销售现金流/营业收入" path="xsxjlyysr">
                  <n-input-number
                    placeholder="请输入销售现金流/营业收入"
                    v-model:value="formValue.xsxjlyysr"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="资产负债比率(%)" path="zcfzl">
                  <n-input-number
                    placeholder="请输入资产负债比率(%)"
                    v-model:value="formValue.zcfzl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="存货周转率(次)" path="chzzl">
                  <n-input-number
                    placeholder="请输入存货周转率(次)"
                    v-model:value="formValue.chzzl"
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
                <n-form-item label="单位: yuan-元" path="unit">
                  <n-input
                    placeholder="请输入单位: yuan-元"
                    v-model:value="formValue.unit"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否为计算指标: 0-原始数据, 1-计算得出" path="isCalculated">
                  <n-input-number
                    placeholder="请输入是否为计算指标: 0-原始数据, 1-计算得出"
                    v-model:value="formValue.isCalculated"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="计算版本" path="calcVersion">
                  <n-input
                    placeholder="请输入计算版本"
                    v-model:value="formValue.calcVersion"
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
  import { Edit, View } from '@/api/financialIndicators';
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