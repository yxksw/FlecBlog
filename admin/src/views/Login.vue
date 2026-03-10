<template>
  <div class="admin-login">
    <div class="login-container">
      <div class="login-content">
        <div class="login-header">
          <h1>后台管理系统</h1>
          <p>Admin Management System</p>
        </div>
        <div class="login-form">
          <div class="form-item">
            <i class="ri-user-line"></i>
            <input type="text" v-model="formState.email" placeholder="请输入邮箱" @keyup.enter="handleLogin" />
          </div>
          <div class="form-item">
            <i class="ri-lock-line"></i>
            <input type="password" v-model="formState.password" placeholder="请输入密码" @keyup.enter="handleLogin" />
          </div>
          <div class="form-submit">
            <button class="submit-btn" :disabled="loading" @click="handleLogin">
              {{ loading ? '登录中...' : '登 录' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { login } from '@/api/user'
import { setTokens } from '@/utils/auth'
import type { LoginParams } from '@/types/user'

const router = useRouter()
const loading = ref(false)
const formState = reactive<LoginParams>({
  email: '',
  password: ''
})

const handleLogin = async () => {
  if (!formState.email || !formState.password) {
    ElMessage.warning('请输入邮箱和密码')
    return
  }

  loading.value = true
  try {
    const { access_token, refresh_token, user } = await login(formState)
    setTokens(access_token, refresh_token)
    // 保存完整用户信息到 localStorage
    localStorage.setItem('userInfo', JSON.stringify(user))
    ElMessage.success('登录成功')
    router.push('/')
  } catch (error) {
    ElMessage.error(error instanceof Error ? error.message : '登录失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
.admin-login {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);

  .login-container {
    width: 100%;
    max-width: 420px;
    padding: 20px;
  }

  .login-content {
    background: rgba(255, 255, 255, 0.95);
    border-radius: 8px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    padding: 40px;
  }

  .login-header {
    text-align: center;
    margin-bottom: 40px;

    h1 {
      font-size: 28px;
      color: #333;
      margin-bottom: 8px;
      font-weight: 500;
    }

    p {
      color: #666;
      font-size: 16px;
    }
  }

  .login-form {
    .form-item {
      position: relative;
      margin-bottom: 24px;

      i {
        position: absolute;
        left: 12px;
        top: 50%;
        transform: translateY(-50%);
        color: #999;
        font-size: 18px;
      }

      input {
        width: 100%;
        height: 42px;
        line-height: 42px;
        padding: 0 15px 0 40px;
        border: 1px solid #dcdfe6;
        border-radius: 4px;
        color: #333;
        transition: all 0.3s ease-in-out;

        &:focus {
          border-color: #667eea;
          box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.2);
        }

        &::placeholder {
          color: #999;
        }
      }
    }

    .form-submit {
      .submit-btn {
        width: 100%;
        height: 42px;
        background: linear-gradient(to right, #667eea, #764ba2);
        border: none;
        border-radius: 4px;
        color: white;
        font-size: 16px;
        cursor: pointer;
        transition: all 0.3s ease-in-out;

        &:hover:not(:disabled) {
          opacity: 0.9;
          transform: translateY(-1px);
        }

        &:disabled {
          opacity: 0.7;
          cursor: not-allowed;
        }
      }
    }
  }

  // 移动端适配
  @media (max-width: 768px) {
    .login-container {
      max-width: 100%;
      padding: 15px;
    }

    .login-content {
      padding: 30px 20px;
    }

    .login-header {
      margin-bottom: 30px;

      h1 {
        font-size: 24px;
      }

      p {
        font-size: 14px;
      }
    }

    .login-form {
      .form-item {
        margin-bottom: 20px;

        input {
          height: 40px;
          font-size: 14px;
        }
      }

      .form-submit {
        .submit-btn {
          height: 40px;
          font-size: 15px;
        }
      }
    }
  }
}
</style>
