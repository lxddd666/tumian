<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑股东户数变化记录表 (记录相邻报告期的户数变化) #' + formValue.id : '添加股东户数变化记录表 (记录相邻报告期的户数变化)'"
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
                <n-form-item label="截止日期 (统计截止日，如2025-12-31)[citation:9]" path="jzrq">
                  <DatePicker v-model:formValue="formValue.jzrq" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="股东户数 (统计截止日的总户数)[citation:3][citation:6]" path="gdhs">
                  <n-input-number
                    placeholder="请输入股东户数 (统计截止日的总户数)[citation:3][citation:6]"
                    v-model:value="formValue.gdhs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]" path="bh">
                  <n-input-number
                    placeholder="请输入比上期变化情况 (通常为百分比，如 -6.23 表示减少6.23%)[citation:9]"
                    v-model:value="formValue.bh"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="变化方向 (衍生字段)" path="changeDirection">
                  <n-input
                    placeholder="请输入变化方向 (衍生字段)"
                    v-model:value="formValue.changeDirection"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
<n-form-item label="数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])" path="dataSource">
                  <n-input
                    placeholder="请输入数据来源 (如: 交易所公告[citation:6]、互动平台[citation:3])"
                    v-model:value="formValue.dataSource"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="公告日期 (信息发布日期)[citation:1]" path="annDate">
                  <DatePicker v-model:formValue="formValue.annDate" type="date" />
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
  import { Edit, View } from '@/api/shareholderChange';
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