<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑ai 选股判断 #' + formValue.id : '添加ai 选股判断'"
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
                <n-form-item label="时间" path="t">
                  <DatePicker v-model:formValue="formValue.t" type="datetime" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="ai id" path="aiId">
                  <n-input-number
                    placeholder="请输入ai id"
                    v-model:value="formValue.aiId"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="指标判断" path="indicatorsJudgment">
<Editor style="height: 450px" id="indicatorsJudgment" v-model:value="formValue.indicatorsJudgment" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="指标判断 1是2否" path="indicatorsFlag">
                  <n-input-number
                    placeholder="请输入指标判断 1是2否"
                    v-model:value="formValue.indicatorsFlag"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="财务判断" path="financialJudgment">
                  <Editor style="height: 450px" id="financialJudgment" v-model:value="formValue.financialJudgment" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="财务判断 1是2否" path="financialFlag">
                  <n-input-number
                    placeholder="请输入财务判断 1是2否"
                    v-model:value="formValue.financialFlag"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="综合判断" path="comprehensiveJudgment">
<Editor style="height: 450px" id="comprehensiveJudgment" v-model:value="formValue.comprehensiveJudgment" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="综合判断 1是2否" path="comprehensiveFlag">
                  <n-input-number
                    placeholder="请输入综合判断 1是2否"
                    v-model:value="formValue.comprehensiveFlag"
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
  import { Edit, View } from '@/api/stockAiJudgment';
  import { State, newState, rules } from './model';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import Editor from '@/components/Editor/editor.vue';
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