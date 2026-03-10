<template>
  <el-form :model="form" label-width="120px" class="setting-form">
    <el-form-item label="网站标题" required>
      <el-input v-model="form.title" placeholder="用于RSS订阅和邮件显示的站点标题" :disabled="loading" />
    </el-form-item>

    <el-form-item label="网站副标题">
      <el-input v-model="form.subtitle" placeholder="网站的副标题或简短描述" :disabled="loading" />
    </el-form-item>

    <el-form-item label="网站作者">
      <el-input v-model="form.author" placeholder="网站作者名称" :disabled="loading" />
    </el-form-item>

    <el-form-item label="作者主页">
      <el-input v-model="form.author_url" placeholder="作者个人主页链接" :disabled="loading" />
    </el-form-item>

    <el-form-item label="作者邮箱">
      <el-input v-model="form.author_email" placeholder="作者联系邮箱" :disabled="loading" />
    </el-form-item>

    <el-form-item label="作者简介">
      <el-input v-model="form.author_desc" placeholder="作者个人简介" :disabled="loading" />
    </el-form-item>

    <el-form-item label="网站标语">
      <el-input v-model="form.slogan" placeholder="网站标语或座右铭" :disabled="loading" />
    </el-form-item>

    <el-form-item label="网站描述">
      <el-input v-model="form.description" type="textarea" :rows="3" placeholder="网站描述，用于SEO" :disabled="loading" />
    </el-form-item>

    <el-form-item label="关键词">
      <el-input v-model="form.keywords" placeholder="网站关键词，多个用逗号分隔" :disabled="loading" />
    </el-form-item>

    <el-form-item label="ICP备案号">
      <el-input v-model="form.icp" placeholder="ICP备案号" :disabled="loading" />
    </el-form-item>

    <el-form-item label="公安备案号">
      <el-input v-model="form.police_record" placeholder="公安备案号" :disabled="loading" />
    </el-form-item>

    <el-form-item label="建站日期">
      <el-date-picker v-model="form.established" type="date" placeholder="选择建站日期" format="YYYY-MM-DD"
        value-format="YYYY-MM-DD" :disabled="loading" style="width: 100%" />
    </el-form-item>

    <el-divider content-position="left">系统地址配置</el-divider>

    <el-form-item label="前台地址" required>
      <el-input v-model="form.web_url" placeholder="例如 https://your-site.com" :disabled="loading" />
    </el-form-item>

    <el-form-item label="后台地址" required>
      <el-input v-model="form.admin_url" placeholder="例如 https://admin.your-site.com" :disabled="loading" />
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
interface SiteForm {
  title: string
  subtitle: string
  author: string
  author_url: string
  author_email: string
  author_desc: string
  slogan: string
  description: string
  keywords: string
  icp: string
  police_record: string
  established: string
  web_url: string
  admin_url: string
}

const form = defineModel<SiteForm>('form', { required: true })

defineProps<{
  loading?: boolean
}>()
</script>

<style lang="scss" scoped>
.setting-form {
  max-width: 800px;
}
</style>
