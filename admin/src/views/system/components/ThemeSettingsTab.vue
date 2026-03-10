<template>
  <el-form :model="form" label-width="120px" class="setting-form">
    <div class="image-row">
      <el-form-item label="网站Favicon">
        <ImageUploader ref="faviconUploaderRef" v-model="form.favicon" upload-type="网站Favicon"
          width="120px" height="120px" />
      </el-form-item>

      <el-form-item label="作者头像">
        <ImageUploader ref="authorAvatarUploaderRef" v-model="form.author_avatar" upload-type="作者头像"
          width="120px" height="120px" />
      </el-form-item>

      <el-form-item label="背景图片">
        <ImageUploader ref="backgroundUploaderRef" v-model="form.background_image" upload-type="背景图片"
          width="213px" height="120px" />
      </el-form-item>

      <el-form-item label="站点截图">
        <ImageUploader ref="screenshotUploaderRef" v-model="form.screenshot" upload-type="站点截图"
          width="213px" height="120px" />
      </el-form-item>
    </div>

    <el-form-item label="打字机文本">
      <JsonListEditor v-model="form.typingTextsList" :fields="typingTextsFields" :default-item="{ value: '' }"
        :disabled="loading" />
    </el-form-item>

    <el-divider content-position="left">社交媒体</el-divider>

    <el-form-item label="侧边栏社交">
      <JsonListEditor v-model="form.sidebarSocialList" :fields="sidebarSocialFields"
        :default-item="{ name: '', url: '', icon: '' }" :disabled="loading" />
    </el-form-item>

    <el-form-item label="页脚社交">
      <JsonListEditor v-model="form.footerSocialList" :fields="footerSocialFields"
        :default-item="{ name: '', url: '', icon: '', position: 'left' }" :disabled="loading" />
    </el-form-item>

    <el-divider content-position="left">关于页面配置</el-divider>

    <el-form-item label="个人描述">
      <el-input v-model="form.about_describe" type="textarea" :rows="3" placeholder="关于页面的个人描述" :disabled="loading" />
    </el-form-item>

    <el-form-item label="描述提示">
      <el-input v-model="form.about_describe_tips" placeholder="例如：前端工程师 · 业余 · 专注 · 享受生活" :disabled="loading" />
    </el-form-item>

    <div class="image-row">
      <el-form-item label="个人照片">
        <ImageUploader ref="aboutPhotoUploaderRef" v-model="form.about_photo" upload-type="个人照片"
          width="120px" height="120px" />
      </el-form-item>

      <el-form-item label="展览图片">
        <ImageUploader ref="aboutExhibitionUploaderRef" v-model="form.about_exhibition" upload-type="展览图片"
          width="213px" height="120px" />
      </el-form-item>
    </div>

    <el-form-item label="个人资料">
      <JsonListEditor v-model="form.profileList" :fields="profileFields" :disabled="loading" hide-controls />
    </el-form-item>

    <el-form-item label="性格类型">
      <el-select v-model="form.about_personality" placeholder="请选择性格类型" clearable :disabled="loading"
        style="width: 100%">
        <el-option-group label="分析家">
          <el-option label="INTJ-A (建筑师-自信型)" value="INTJ-A" />
          <el-option label="INTJ-T (建筑师-波动型)" value="INTJ-T" />
          <el-option label="INTP-A (逻辑学家-自信型)" value="INTP-A" />
          <el-option label="INTP-T (逻辑学家-波动型)" value="INTP-T" />
          <el-option label="ENTJ-A (指挥官-自信型)" value="ENTJ-A" />
          <el-option label="ENTJ-T (指挥官-波动型)" value="ENTJ-T" />
          <el-option label="ENTP-A (辩论家-自信型)" value="ENTP-A" />
          <el-option label="ENTP-T (辩论家-波动型)" value="ENTP-T" />
        </el-option-group>
        <el-option-group label="外交家">
          <el-option label="INFJ-A (提倡者-自信型)" value="INFJ-A" />
          <el-option label="INFJ-T (提倡者-波动型)" value="INFJ-T" />
          <el-option label="INFP-A (调停者-自信型)" value="INFP-A" />
          <el-option label="INFP-T (调停者-波动型)" value="INFP-T" />
          <el-option label="ENFJ-A (主人公-自信型)" value="ENFJ-A" />
          <el-option label="ENFJ-T (主人公-波动型)" value="ENFJ-T" />
          <el-option label="ENFP-A (竞选者-自信型)" value="ENFP-A" />
          <el-option label="ENFP-T (竞选者-波动型)" value="ENFP-T" />
        </el-option-group>
        <el-option-group label="守护者">
          <el-option label="ISTJ-A (物流师-自信型)" value="ISTJ-A" />
          <el-option label="ISTJ-T (物流师-波动型)" value="ISTJ-T" />
          <el-option label="ISFJ-A (守卫者-自信型)" value="ISFJ-A" />
          <el-option label="ISFJ-T (守卫者-波动型)" value="ISFJ-T" />
          <el-option label="ESTJ-A (总经理-自信型)" value="ESTJ-A" />
          <el-option label="ESTJ-T (总经理-波动型)" value="ESTJ-T" />
          <el-option label="ESFJ-A (执政官-自信型)" value="ESFJ-A" />
          <el-option label="ESFJ-T (执政官-波动型)" value="ESFJ-T" />
        </el-option-group>
        <el-option-group label="探险家">
          <el-option label="ISTP-A (鉴赏家-自信型)" value="ISTP-A" />
          <el-option label="ISTP-T (鉴赏家-波动型)" value="ISTP-T" />
          <el-option label="ISFP-A (探险家-自信型)" value="ISFP-A" />
          <el-option label="ISFP-T (探险家-波动型)" value="ISFP-T" />
          <el-option label="ESTP-A (企业家-自信型)" value="ESTP-A" />
          <el-option label="ESTP-T (企业家-波动型)" value="ESTP-T" />
          <el-option label="ESFP-A (表演者-自信型)" value="ESFP-A" />
          <el-option label="ESFP-T (表演者-波动型)" value="ESFP-T" />
        </el-option-group>
      </el-select>
    </el-form-item>

    <el-form-item label="座右铭">
      <div class="motto-inputs">
        <el-input v-model="form.mottoMainList[0]" placeholder="第 1 行" :disabled="loading" />
        <el-input v-model="form.mottoMainList[1]" placeholder="第 2 行" :disabled="loading" />
      </div>
    </el-form-item>

    <el-form-item label="一言">
      <el-input v-model="form.about_motto_sub" placeholder="一句话介绍" :disabled="loading" />
    </el-form-item>

    <el-form-item label="联系方式">
      <JsonListEditor v-model="form.socializeList" :fields="socializeFields" :default-item="{ name: '', url: '' }"
        :disabled="loading" />
    </el-form-item>

    <el-form-item label="创作平台">
      <JsonListEditor v-model="form.creationList" :fields="creationFields" :default-item="{ name: '', url: '' }"
        :disabled="loading" />
    </el-form-item>

    <el-form-item label="版本信息">
      <JsonListEditor v-model="form.versionsList" :fields="versionsFields" :disabled="loading" hide-controls />
    </el-form-item>

    <el-form-item label="站长联盟">
      <JsonListEditor v-model="form.unionsList" :fields="unionsFields" :default-item="{ name: '', url: '' }"
        :disabled="loading" />
    </el-form-item>

    <el-form-item label="心路历程">
      <el-input v-model="form.about_story" type="textarea" :rows="6" placeholder="关于本站的介绍和心路历程" :disabled="loading" />
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ImageUploader from '@/components/common/ImageUploader.vue'
import JsonListEditor from '@/components/common/JsonListEditor.vue'
import type { FieldConfig } from '@/components/common/JsonListEditor.vue'

interface ThemeFormData {
  author_avatar: string
  favicon: string
  background_image: string
  screenshot: string
  about_describe: string
  about_describe_tips: string
  about_photo: string
  about_exhibition: string
  about_personality: string
  about_motto_sub: string
  about_story: string
  typingTextsList: Array<{ value: string }>
  sidebarSocialList: Array<{ name: string; url: string; icon: string }>
  footerSocialList: Array<{ name: string; url: string; icon: string; position: string }>
  profileList: Array<{ label: string; value: string; color: string }>
  mottoMainList: string[]
  socializeList: Array<{ name: string; url: string }>
  creationList: Array<{ name: string; url: string }>
  versionsList: Array<{ name: string; version: string }>
  unionsList: Array<{ name: string; url: string }>
}

const form = defineModel<ThemeFormData>('form', { required: true })

defineProps<{
  loading?: boolean
}>()

// 图片上传器引用
const authorAvatarUploaderRef = ref<InstanceType<typeof ImageUploader>>()
const faviconUploaderRef = ref<InstanceType<typeof ImageUploader>>()
const backgroundUploaderRef = ref<InstanceType<typeof ImageUploader>>()
const screenshotUploaderRef = ref<InstanceType<typeof ImageUploader>>()
const aboutPhotoUploaderRef = ref<InstanceType<typeof ImageUploader>>()
const aboutExhibitionUploaderRef = ref<InstanceType<typeof ImageUploader>>()

// 预设的常用社交平台图标
const commonIcons = [
  'github-line', 'mail-line', 'twitter-x-line', 'bilibili-line', 'wechat-line',
  'qq-line', 'weibo-line', 'zhihu-line', 'douban-line', 'linkedin-line',
  'facebook-line', 'instagram-line', 'youtube-line', 'tiktok-line', 'discord-line',
  'telegram-line', 'slack-line', 'rss-line', 'links-line'
]

const iconOptions = commonIcons.map(icon => ({ label: icon, value: icon, icon: 'ri-' + icon }))
const iconField = {
  key: 'icon',
  type: 'select' as const,
  placeholder: '图标',
  style: 'width: 200px; margin-right: 8px',
  prefix: 'ri-',
  filterable: true,
  allowCreate: true,
  options: iconOptions
}

// 字段配置
const typingTextsFields: FieldConfig[] = [
  { key: 'value', type: 'text', placeholder: '打字机文本', style: 'flex: 1' }
]

const sidebarSocialFields: FieldConfig[] = [
  { key: 'name', type: 'text', placeholder: '平台名称', style: 'width: 120px' },
  { key: 'url', type: 'text', placeholder: '链接地址', style: 'flex: 1; margin: 0 8px' },
  iconField
]

const footerSocialFields: FieldConfig[] = [
  { key: 'name', type: 'text', placeholder: '平台名称', style: 'width: 100px' },
  { key: 'url', type: 'text', placeholder: '链接地址', style: 'flex: 1; margin: 0 8px' },
  iconField,
  {
    key: 'position',
    type: 'select',
    placeholder: '位置',
    style: 'width: 80px; margin-right: 8px',
    options: [
      { label: '左', value: 'left' },
      { label: '右', value: 'right' }
    ]
  }
]

const nameUrlFields: FieldConfig[] = [
  { key: 'name', type: 'text', placeholder: '平台名称', style: 'width: 120px' },
  { key: 'url', type: 'text', placeholder: '链接地址', style: 'flex: 1; margin: 0 8px' }
]

const socializeFields = nameUrlFields
const creationFields = nameUrlFields

const unionsFields: FieldConfig[] = [
  { key: 'name', type: 'text', placeholder: '联盟名称', style: 'width: 150px' },
  { key: 'url', type: 'text', placeholder: '链接地址', style: 'flex: 1; margin: 0 8px' }
]

const profileFields: FieldConfig[] = [
  { key: 'label', type: 'text', placeholder: '标签', style: 'width: 100px' },
  { key: 'value', type: 'text', placeholder: '值', style: 'flex: 1; margin: 0 8px' },
  { key: 'color', type: 'color' }
]

const versionsFields: FieldConfig[] = [
  { key: 'name', type: 'text', placeholder: '技术名称', style: 'width: 150px' },
  { key: 'version', type: 'text', placeholder: '版本号', style: 'flex: 1; margin: 0 8px' }
]

// 暴露上传器引用给父组件
defineExpose({
  authorAvatarUploaderRef,
  faviconUploaderRef,
  backgroundUploaderRef,
  screenshotUploaderRef,
  aboutPhotoUploaderRef,
  aboutExhibitionUploaderRef
})
</script>

<style lang="scss" scoped>
.setting-form {
  max-width: 800px;

  .image-row {
    display: flex;
    gap: 40px;

    .el-form-item {
      margin-bottom: 22px;
    }
  }
}

.motto-inputs {
  display: flex;
  gap: 8px;
}
</style>
