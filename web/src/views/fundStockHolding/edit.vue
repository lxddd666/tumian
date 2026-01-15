<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑基金持股明细表 (来源于基金定期报告) #' + formValue.id : '添加基金持股明细表 (来源于基金定期报告)'"
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
                <n-form-item label="截止日期 (报告期，如2025-12-31)" path="jzrq">
                  <DatePicker v-model:formValue="formValue.jzrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易时间 (衍生自jzrq，兼容时间序列查询)" path="t">
                  <DatePicker v-model:formValue="formValue.t" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="基金名称" path="jjmc">
                  <n-input
                    placeholder="请输入基金名称"
                    v-model:value="formValue.jjmc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="基金代码" path="jjdm">
                  <n-input
                    placeholder="请输入基金代码"
                    v-model:value="formValue.jjdm"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="股票代码 (如: 000001.SZ)" path="symbol">
                  <n-input
                    placeholder="请输入股票代码 (如: 000001.SZ)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="持仓数量(股)" path="ccsl">
                  <n-input-number
                    placeholder="请输入持仓数量(股)"
                    v-model:value="formValue.ccsl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="占流通股比例(%)" path="ltbl">
                  <n-input-number
                    placeholder="请输入占流通股比例(%)"
                    v-model:value="formValue.ltbl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="持股市值（元）" path="cgsz">
                  <n-input-number
                    placeholder="请输入持股市值（元）"
                    v-model:value="formValue.cgsz"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="占净值比例（%）" path="jzbl">
                  <n-input-number
                    placeholder="请输入占净值比例（%）"
                    v-model:value="formValue.jzbl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="估算持仓成本 (元/股)" path="avgCost">
                  <n-input-number
                    placeholder="请输入估算持仓成本 (元/股)"
                    v-model:value="formValue.avgCost"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据来源 (如: 基金季报)" path="dataSource">
                  <n-input
                    placeholder="请输入数据来源 (如: 基金季报)"
                    v-model:value="formValue.dataSource"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="报告类型: 季报, 中报, 年报" path="reportType">
                  <n-input
                    placeholder="请输入报告类型: 季报, 中报, 年报"
                    v-model:value="formValue.reportType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否为该基金对该股票的最新持仓: 0-否, 1-是" path="isLatest">
                  <n-input-number
                    placeholder="请输入是否为该基金对该股票的最新持仓: 0-否, 1-是"
                    v-model:value="formValue.isLatest"
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
  import { Edit, View } from '@/api/fundStockHolding';
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