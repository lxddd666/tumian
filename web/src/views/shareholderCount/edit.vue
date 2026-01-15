<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑公司股东户数统计表 (按报告期统计) #' + formValue.id : '添加公司股东户数统计表 (按报告期统计)'"
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
                <n-form-item label="公司代码/股票代码 (例如: 000001.SZ, 600000.SS)" path="symbol">
                  <n-input
                    placeholder="请输入公司代码/股票代码 (例如: 000001.SZ, 600000.SS)"
                    v-model:value="formValue.symbol"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="截止日期 (报告期结束日, 如2023-09-30)" path="jzrq">
                  <DatePicker v-model:formValue="formValue.jzrq" type="date" />
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
<n-form-item label="报告类型: annual-年报, half_year-中报, quarter-季报" path="reportType">
                  <n-input
                    placeholder="请输入报告类型: annual-年报, half_year-中报, quarter-季报"
                    v-model:value="formValue.reportType"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="股东总数 (户)" path="gdzs">
                  <n-input-number
                    placeholder="请输入股东总数 (户)"
                    v-model:value="formValue.gdzs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="A股东户数 (户)" path="agdhs">
                  <n-input-number
                    placeholder="请输入A股东户数 (户)"
                    v-model:value="formValue.agdhs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="B股东户数 (户)" path="bgdhs">
                  <n-input-number
                    placeholder="请输入B股东户数 (户)"
                    v-model:value="formValue.bgdhs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="H股东户数 (户)" path="hgdhs">
                  <n-input-number
                    placeholder="请输入H股东户数 (户)"
                    v-model:value="formValue.hgdhs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="已流通股东户数 (户)" path="yltgdhs">
                  <n-input-number
                    placeholder="请输入已流通股东户数 (户)"
                    v-model:value="formValue.yltgdhs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="未流通股东户数 (户)" path="wltgdhs">
                  <n-input-number
                    placeholder="请输入未流通股东户数 (户)"
                    v-model:value="formValue.wltgdhs"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="A股股东占比(%)" path="agRatio">
                  <n-input-number
                    placeholder="请输入A股股东占比(%)"
                    v-model:value="formValue.agRatio"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="已流通股东占比(%)" path="yltRatio">
                  <n-input-number
                    placeholder="请输入已流通股东占比(%)"
                    v-model:value="formValue.yltRatio"
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
                <n-form-item label="是否为该报告期最新数据" path="isLatest">
                  <n-input-number
                    placeholder="请输入是否为该报告期最新数据"
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
  import { Edit, View } from '@/api/shareholderCount';
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