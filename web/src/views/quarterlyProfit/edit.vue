<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑季度利润数据表 (近一年各季度) #' + formValue.id : '添加季度利润数据表 (近一年各季度)'"
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
                <n-form-item label="截止日期 (报告期截止日，如2025-03-31)" path="date">
                  <DatePicker v-model:formValue="formValue.date" type="date" />
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
                <n-form-item label="会计期间 (如2025Q1)" path="fiscalPeriod">
                  <n-input
                    placeholder="请输入会计期间 (如2025Q1)"
                    v-model:value="formValue.fiscalPeriod"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业收入（万元）" path="income">
                  <n-input-number
                    placeholder="请输入营业收入（万元）"
                    v-model:value="formValue.income"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业支出（万元）" path="expend">
                  <n-input-number
                    placeholder="请输入营业支出（万元）"
                    v-model:value="formValue.expend"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业利润（万元）" path="profit">
                  <n-input-number
                    placeholder="请输入营业利润（万元）"
                    v-model:value="formValue.profit"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="利润总额（万元）" path="totalp">
                  <n-input-number
                    placeholder="请输入利润总额（万元）"
                    v-model:value="formValue.totalp"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利润（万元）" path="reprofit">
                  <n-input-number
                    placeholder="请输入净利润（万元）"
                    v-model:value="formValue.reprofit"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="基本每股收益(元/股)" path="basege">
                  <n-input-number
                    placeholder="请输入基本每股收益(元/股)"
                    v-model:value="formValue.basege"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="稀释每股收益(元/股)" path="ettege">
                  <n-input-number
                    placeholder="请输入稀释每股收益(元/股)"
                    v-model:value="formValue.ettege"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="其他综合收益（万元）" path="otherp">
                  <n-input-number
                    placeholder="请输入其他综合收益（万元）"
                    v-model:value="formValue.otherp"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="综合收益总额（万元）" path="totalcp">
                  <n-input-number
                    placeholder="请输入综合收益总额（万元）"
                    v-model:value="formValue.totalcp"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="毛利率(%)" path="grossProfitMargin">
                  <n-input-number
                    placeholder="请输入毛利率(%)"
                    v-model:value="formValue.grossProfitMargin"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="净利率(%)" path="netProfitMargin">
                  <n-input-number
                    placeholder="请输入净利率(%)"
                    v-model:value="formValue.netProfitMargin"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="营业利润率(%)" path="operatingProfitRatio">
                  <n-input-number
                    placeholder="请输入营业利润率(%)"
                    v-model:value="formValue.operatingProfitRatio"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告类型: 一季报, 中报, 三季报, 年报" path="reportType">
                  <n-input
                    placeholder="请输入报告类型: 一季报, 中报, 三季报, 年报"
                    v-model:value="formValue.reportType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据来源 (如: 交易所财报)" path="dataSource">
                  <n-input
                    placeholder="请输入数据来源 (如: 交易所财报)"
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
                <n-form-item label="是否审计: 0-未审计, 1-已审计" path="isAudited">
                  <n-input-number
                    placeholder="请输入是否审计: 0-未审计, 1-已审计"
                    v-model:value="formValue.isAudited"
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
  import { Edit, View } from '@/api/quarterlyProfit';
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