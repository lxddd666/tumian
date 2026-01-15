<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑移动平均线(MA)指标数据表 #' + formValue.id : '添加移动平均线(MA)指标数据表'"
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
                <n-form-item label="股票或标的代码 (例如: AAPL, 000001.SZ)" path="symbol">
                  <n-input
                    placeholder="请输入股票或标的代码 (例如: AAPL, 000001.SZ)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="交易时间 (统一为datetime类型，日线数据时间部分设为00:00:00)" path="t">
                  <DatePicker v-model:formValue="formValue.t" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="数据间隔: minute-短分时, day-日线" path="intervalType">
                  <n-input
                    placeholder="请输入数据间隔: minute-短分时, day-日线"
                    v-model:value="formValue.intervalType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA3值" path="ma3">
                  <n-input-number
                    placeholder="请输入MA3值"
                    v-model:value="formValue.ma3"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA5值" path="ma5">
                  <n-input-number
                    placeholder="请输入MA5值"
                    v-model:value="formValue.ma5"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA10值" path="ma10">
                  <n-input-number
                    placeholder="请输入MA10值"
                    v-model:value="formValue.ma10"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA15值" path="ma15">
                  <n-input-number
                    placeholder="请输入MA15值"
                    v-model:value="formValue.ma15"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA20值" path="ma20">
                  <n-input-number
                    placeholder="请输入MA20值"
                    v-model:value="formValue.ma20"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA30值" path="ma30">
                  <n-input-number
                    placeholder="请输入MA30值"
                    v-model:value="formValue.ma30"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA60值" path="ma60">
                  <n-input-number
                    placeholder="请输入MA60值"
                    v-model:value="formValue.ma60"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA120值" path="ma120">
                  <n-input-number
                    placeholder="请输入MA120值"
                    v-model:value="formValue.ma120"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA200值" path="ma200">
                  <n-input-number
                    placeholder="请输入MA200值"
                    v-model:value="formValue.ma200"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="MA250值" path="ma250">
                  <n-input-number
                    placeholder="请输入MA250值"
                    v-model:value="formValue.ma250"
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
  import { Edit, View } from '@/api/maData';
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