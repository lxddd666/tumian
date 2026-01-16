<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑资金流向明细表 #' + formValue.id : '添加资金流向明细表'"
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
                <n-form-item label="交易时间 (通常为HHMMSS格式的整数)" path="t">
                  <n-input-number
                    placeholder="请输入交易时间 (通常为HHMMSS格式的整数)"
                    v-model:value="formValue.t"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买单总单数" path="zmbzds">
                  <n-input-number
                    placeholder="请输入主买单总单数"
                    v-model:value="formValue.zmbzds"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖单总单数" path="zmszds">
                  <n-input-number
                    placeholder="请输入主卖单总单数"
                    v-model:value="formValue.zmszds"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="大单动向" path="dddx">
                  <n-input-number
                    placeholder="请输入大单动向"
                    v-model:value="formValue.dddx"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="涨跌动因" path="zddy">
                  <n-input-number
                    placeholder="请输入涨跌动因"
                    v-model:value="formValue.zddy"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="大单差分" path="ddcf">
                  <n-input-number
                    placeholder="请输入大单差分"
                    v-model:value="formValue.ddcf"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买单总单数增量" path="zmbzdszl">
                  <n-input-number
                    placeholder="请输入主买单总单数增量"
                    v-model:value="formValue.zmbzdszl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖单总单数增量" path="zmszdszl">
                  <n-input-number
                    placeholder="请输入主卖单总单数增量"
                    v-model:value="formValue.zmszdszl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="成交笔数增量" path="cjbszl">
                  <n-input-number
                    placeholder="请输入成交笔数增量"
                    v-model:value="formValue.cjbszl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买特大单成交额" path="zmbtdcje">
                  <n-input-number
                    placeholder="请输入主买特大单成交额"
                    v-model:value="formValue.zmbtdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买大单成交额" path="zmbddcje">
                  <n-input-number
                    placeholder="请输入主买大单成交额"
                    v-model:value="formValue.zmbddcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买中单成交额" path="zmbzdcje">
                  <n-input-number
                    placeholder="请输入主买中单成交额"
                    v-model:value="formValue.zmbzdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买小单成交额" path="zmbxdcje">
                  <n-input-number
                    placeholder="请输入主买小单成交额"
                    v-model:value="formValue.zmbxdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖特大单成交额" path="zmstdcje">
                  <n-input-number
                    placeholder="请输入主卖特大单成交额"
                    v-model:value="formValue.zmstdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖大单成交额" path="zmsddcje">
                  <n-input-number
                    placeholder="请输入主卖大单成交额"
                    v-model:value="formValue.zmsddcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖中单成交额" path="zmszdcje">
                  <n-input-number
                    placeholder="请输入主卖中单成交额"
                    v-model:value="formValue.zmszdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖小单成交额" path="zmsxdcje">
                  <n-input-number
                    placeholder="请输入主卖小单成交额"
                    v-model:value="formValue.zmsxdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买特大单成交额" path="bdmbtdcje">
                  <n-input-number
                    placeholder="请输入被动买特大单成交额"
                    v-model:value="formValue.bdmbtdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买大单成交额" path="bdmbddcje">
                  <n-input-number
                    placeholder="请输入被动买大单成交额"
                    v-model:value="formValue.bdmbddcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买中单成交额" path="bdmbzdcje">
                  <n-input-number
                    placeholder="请输入被动买中单成交额"
                    v-model:value="formValue.bdmbzdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买小单成交额" path="bdmbxdcje">
                  <n-input-number
                    placeholder="请输入被动买小单成交额"
                    v-model:value="formValue.bdmbxdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖特大单成交额" path="bdmstdcje">
                  <n-input-number
                    placeholder="请输入被动卖特大单成交额"
                    v-model:value="formValue.bdmstdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖大单成交额" path="bdmsddcje">
                  <n-input-number
                    placeholder="请输入被动卖大单成交额"
                    v-model:value="formValue.bdmsddcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖中单成交额" path="bdmszdcje">
                  <n-input-number
                    placeholder="请输入被动卖中单成交额"
                    v-model:value="formValue.bdmszdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖小单成交额" path="bdmsxdcje">
                  <n-input-number
                    placeholder="请输入被动卖小单成交额"
                    v-model:value="formValue.bdmsxdcje"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买特大单成交量" path="zmbtdcjl">
                  <n-input-number
                    placeholder="请输入主买特大单成交量"
                    v-model:value="formValue.zmbtdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买大单成交量" path="zmbddcjl">
                  <n-input-number
                    placeholder="请输入主买大单成交量"
                    v-model:value="formValue.zmbddcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买中单成交量" path="zmbzdcjl">
                  <n-input-number
                    placeholder="请输入主买中单成交量"
                    v-model:value="formValue.zmbzdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买小单成交量" path="zmbxdcjl">
                  <n-input-number
                    placeholder="请输入主买小单成交量"
                    v-model:value="formValue.zmbxdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖特大单成交量" path="zmstdcjl">
                  <n-input-number
                    placeholder="请输入主卖特大单成交量"
                    v-model:value="formValue.zmstdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖大单成交量" path="zmsddcjl">
                  <n-input-number
                    placeholder="请输入主卖大单成交量"
                    v-model:value="formValue.zmsddcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖中单成交量" path="zmszdcjl">
                  <n-input-number
                    placeholder="请输入主卖中单成交量"
                    v-model:value="formValue.zmszdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主卖小单成交量" path="zmsxdcjl">
                  <n-input-number
                    placeholder="请输入主卖小单成交量"
                    v-model:value="formValue.zmsxdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买特大单成交量" path="bdmbtdcjl">
                  <n-input-number
                    placeholder="请输入被动买特大单成交量"
                    v-model:value="formValue.bdmbtdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买大单成交量" path="bdmbddcjl">
                  <n-input-number
                    placeholder="请输入被动买大单成交量"
                    v-model:value="formValue.bdmbddcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买中单成交量" path="bdmbzdcjl">
                  <n-input-number
                    placeholder="请输入被动买中单成交量"
                    v-model:value="formValue.bdmbzdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动买小单成交量" path="bdmbxdcjl">
                  <n-input-number
                    placeholder="请输入被动买小单成交量"
                    v-model:value="formValue.bdmbxdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖特大单成交量" path="bdmstdcjl">
                  <n-input-number
                    placeholder="请输入被动卖特大单成交量"
                    v-model:value="formValue.bdmstdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖大单成交量" path="bdmsddcjl">
                  <n-input-number
                    placeholder="请输入被动卖大单成交量"
                    v-model:value="formValue.bdmsddcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖中单成交量" path="bdmszdcjl">
                  <n-input-number
                    placeholder="请输入被动卖中单成交量"
                    v-model:value="formValue.bdmszdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="被动卖小单成交量" path="bdmsxdcjl">
                  <n-input-number
                    placeholder="请输入被动卖小单成交量"
                    v-model:value="formValue.bdmsxdcjl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买特大单成交额增量" path="zmbtdcjzl">
                  <n-input-number
                    placeholder="请输入主买特大单成交额增量"
                    v-model:value="formValue.zmbtdcjzl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买大单成交额增量" path="zmbddcjzl">
                  <n-input-number
                    placeholder="请输入主买大单成交额增量"
                    v-model:value="formValue.zmbddcjzl"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买特大单成交量增量" path="zmbtdcjzlv">
                  <n-input-number
                    placeholder="请输入主买特大单成交量增量"
                    v-model:value="formValue.zmbtdcjzlv"
                    />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主买大单成交量增量" path="zmbddcjzlv">
                  <n-input-number
                    placeholder="请输入主买大单成交量增量"
                    v-model:value="formValue.zmbddcjzlv"
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
  import { Edit, View } from '@/api/flowOfFunds';
  import { State, newState, rules } from './model';
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