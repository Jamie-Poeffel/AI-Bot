<!-- LoginView.vue -->
<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'

/* ---------- Login-State ---------- */
const email = ref('')
const password = ref('')
const errorMsg = ref('')
const loading = ref(false)
const router = useRouter()

async function handleLogin() {
  loading.value = true
  try {
    const resp = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        Email: email.value.toLowerCase(),
        Password: password.value,
      }),
      credentials: 'include',
    });


    if (!resp.ok) {
      const { message } = await resp.json().catch(() => ({}))
      throw new Error(message || 'Ungültige Anmeldedaten')
    }

    router.push('/home');
  } catch (e: any) {
    errorMsg.value = e.message
  } finally {
    loading.value = false
  }
}

/* ---------- Theme-State ---------- */
const theme = ref<'light' | 'dark'>('light')

function applyTheme(t: 'light' | 'dark') {
  document.documentElement.setAttribute('data-theme', t)
  localStorage.setItem('theme', t)
}

onMounted(() => {
  const saved = localStorage.getItem('theme') as 'light' | 'dark' | null
  if (saved) theme.value = saved
  else if (window.matchMedia('(prefers-color-scheme: dark)').matches) theme.value = 'dark'
  applyTheme(theme.value)
})

watch(theme, (t) => applyTheme(t))
</script>

<template>
  <!-- Theme-Auswahl oben rechts -->
  <div class="theme-selector">
    <select v-model="theme">
      <option value="light">Hell</option>
      <option value="dark">Dunkel</option>
    </select>
  </div>

  <!-- Login-Card -->
  <div class="login-page">
    <div class="login-card">
      <h1>Login</h1>

      <form @submit.prevent="handleLogin">
        <label for="email">E-Mail</label>
        <input id="email" v-model="email" type="email" placeholder="name@example.com" autocomplete="username"
          required />

        <label for="password">Passwort</label>

        <input id="password" v-model="password" type="password" placeholder="••••••••" autocomplete="current-password"
          required />

        <button :disabled="loading">
          {{ loading ? 'Anmelden …' : 'Anmelden' }}
        </button>

        <router-link class="router-link" to="/register">Noch kein Konto? Jetzt registrieren</router-link>

        <p v-if="errorMsg" class="error">{{ errorMsg }}</p>
      </form>
    </div>
  </div>
</template>

<!-- ---------- GLOBAL: Farb-Variablen & Dark-Override ---------- -->
<style></style>

<!-- ---------- COMPONENT-STYLES (scoped) ---------- -->
<style scoped>
.router-link {
  display: block;
  margin-top: .5rem;
  text-align: center;
  color: var(--primary);
  text-decoration: none;
  font-size: 1.1rem;
}
.router-link {
  display: block;
  margin-top: .5rem;
  text-align: center;
  color: var(--primary);
  text-decoration: none;
  font-size: 1.1rem;
}
.theme-selector {
  position: fixed;
  top: 1rem;
  right: 1rem;
}

.theme-selector select {
  padding: .4rem .75rem;
  border-radius: .5rem;
  border: 1px solid var(--input-border);
  background: var(--input-bg);
  color: var(--text);
  font-size: .875rem;
}

.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg);
  font-family: system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif;
  color: var(--text);
  padding: 2rem;
}

.login-card {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  border-radius: 1.5rem;
  background: var(--card-bg);
  backdrop-filter: blur(8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, .2);
}

.login-card h1 {
  margin: 0 0 1.5rem;
  text-align: center;
  font-size: 2rem;
  font-weight: 700;
}

.login-card form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.login-card label {
  font-size: .875rem;
  font-weight: 500;
  color: var(--muted);
}

.login-card input {
  width: 100%;
  padding: .5rem 1rem;
  border: 2px solid var(--input-border);
  border-radius: .75rem;
  background: var(--input-bg);
  color: var(--text);
  font-size: 1rem;
  transition: border-color .15s, box-shadow .15s;
}

.login-card input::placeholder {
  color: var(--muted);
}

.login-card input:focus {
  outline: none;
  border-color: var(--border-focus);
  box-shadow: 0 0 0 2px var(--border-focus);
}

.login-card button {
  padding: .75rem;
  border: none;
  border-radius: .75rem;
  background: var(--primary);
  color: #fff;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background .15s, opacity .15s;
}

.login-card button:hover {
  background: var(--primary-hov);
}

.login-card button:disabled {
  opacity: .5;
  cursor: not-allowed;
}

.error {
  text-align: center;
  font-size: .875rem;
  color: var(--error);
}

html,
body {
  height: 100%;
  width: 100%;
  margin: 0;
  /* weiße Standardränder eliminieren */
  background: var(--bg);
  /* damit die Hintergrundfarbe immer das ganze Fenster ausfüllt */
}

body {
  color: var(--hintergrund);
}
</style>
