<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑股票基础信息表 #' + formValue.id : '添加股票基础信息表'"
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
                <n-form-item label="股票代码" path="symbol">
                  <n-input
                    placeholder="请输入股票代码"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="ii" path="ii">
                  <n-input
                    placeholder="请输入ii"
                    v-model:value="formValue.ii"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="ei" path="ei">
                  <n-input
                    placeholder="请输入ei"
                    v-model:value="formValue.ei"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易所名称" path="exchange">
                  <n-input
                    placeholder="请输入交易所名称"
                    v-model:value="formValue.exchange"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="name" path="name">
                  <n-input
                    placeholder="请输入name"
                    v-model:value="formValue.name"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="short_name" path="shortName">
                  <n-input
                    placeholder="请输入short_name"
                    v-model:value="formValue.shortName"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="en_name" path="enName">
                  <n-input
                    placeholder="请输入en_name"
                    v-model:value="formValue.enName"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="上市日期" path="od">
                  <DatePicker v-model:formValue="formValue.od" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="data_update_date" path="dataUpdateDate">
                  <DatePicker v-model:formValue="formValue.dataUpdateDate" type="date" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="pc" path="pc">
                  <n-input-number
                    placeholder="请输入pc"
                    v-model:value="formValue.pc"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="up" path="up">
                  <n-input-number
                    placeholder="请输入up"
                    v-model:value="formValue.up"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="dp" path="dp">
                  <n-input-number
                    placeholder="请输入dp"
                    v-model:value="formValue.dp"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="pk" path="pk">
                  <n-input-number
                    placeholder="请输入pk"
                    v-model:value="formValue.pk"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="fv" path="fv">
                  <n-input-number
                    placeholder="请输入fv"
                    v-model:value="formValue.fv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="tv" path="tv">
                  <n-input-number
                    placeholder="请输入tv"
                    v-model:value="formValue.tv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="流通股比例 (%)" path="floatRatio">
                  <n-input-number
                    placeholder="请输入流通股比例 (%)"
                    v-model:value="formValue.floatRatio"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="is" path="is">
                  <n-input-number
                    placeholder="请输入is"
                    v-model:value="formValue.is"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="交易状态描述" path="tradingStatus">
                  <n-input
                    placeholder="请输入交易状态描述"
                    v-model:value="formValue.tradingStatus"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="industry" path="industry">
                  <n-input
                    placeholder="请输入industry"
                    v-model:value="formValue.industry"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="sector" path="sector">
                  <n-input
                    placeholder="请输入sector"
                    v-model:value="formValue.sector"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="market_type" path="marketType">
                  <n-input
                    placeholder="请输入market_type"
                    v-model:value="formValue.marketType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="data_source" path="dataSource">
                  <n-input
                    placeholder="请输入data_source"
                    v-model:value="formValue.dataSource"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="is_active" path="isActive">
                  <n-input-number
                    placeholder="请输入is_active"
                    v-model:value="formValue.isActive"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="version" path="version">
                  <n-input-number
                    placeholder="请输入version"
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
  import { Edit, View } from '@/api/stockBasicInfo';
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