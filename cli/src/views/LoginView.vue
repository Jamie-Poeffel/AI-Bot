<!-- LoginView.vue -->
<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'

/* ---------- Login-State ---------- */
const email    = ref('')
const password = ref('')
const errorMsg = ref('')
const loading  = ref(false)
const router   = useRouter()

async function handleLogin () {
  loading.value = true
  try {
    const resp = await fetch('http://localhost:3000/login', {
      method : 'POST',
      headers: { 'Content-Type': 'application/json' },
      body   : JSON.stringify({ email: email.value, password: password.value })
    })

    if (!resp.ok) {
      const { message } = await resp.json().catch(() => ({}))
      throw new Error(message || 'Ungültige Anmeldedaten')
    }

    const { token } = await resp.json()
    localStorage.setItem('token', token)
    router.push('/')                // Ziel anpassen
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
        <input id="email" v-model="email" type="email"
               placeholder="name@example.com" autocomplete="username" required />

        <label for="password">Passwort</label>
        <input id="password" v-model="password" type="password"
               placeholder="••••••••" autocomplete="current-password" required />

        <button :disabled="loading">
          {{ loading ? 'Anmelden …' : 'Anmelden' }}
        </button>

        <p v-if="errorMsg" class="error">{{ errorMsg }}</p>
      </form>
    </div>
  </div>
</template>

<!-- ---------- GLOBAL: Farb-Variablen & Dark-Override ---------- -->
<style>
:root {
  --bg            : #f0f4ff;
  --card-bg       : rgba(255,255,255,0.95);
  --input-bg      : #ffffff;
  --input-border  : #d1d5db;
  --border-focus  : #4f46e5;
  --text          : #0f172a;
  --muted         : #64748b;
  --primary       : #4f46e5;
  --primary-hov   : #4338ca;
  --error         : #ef4444;
}

html[data-theme="dark"] {
  --bg            : #0f172a;
  --card-bg       : rgba(30,41,59,0.82);
  --input-bg      : rgba(51,65,85,0.60);
  --input-border  : #64748b;
  --text          : #e2e8f0;
  --muted         : #94a3b8;
}
</style>

<!-- ---------- COMPONENT-STYLES (scoped) ---------- -->
<style scoped>
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
  font-family: system-ui,-apple-system,'Segoe UI',Roboto,sans-serif;
  color: var(--text);
  padding: 2rem;
}

.login-card {
  width: 100%;
  max-width: 24rem;
  padding: 2rem;
  border-radius: 1.5rem;
  background: var(--card-bg);
  backdrop-filter: blur(8px);
  box-shadow: 0 20px 40px rgba(0,0,0,.2);
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

.login-card input::placeholder { color: var(--muted); }

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

.login-card button:hover    { background: var(--primary-hov); }
.login-card button:disabled { opacity: .5; cursor: not-allowed; }

.error {
  text-align: center;
  font-size: .875rem;
  color: var(--error);
}

html, body {
  height : 100%;
  width  : 100%;
  margin : 0;          /* weiße Standardränder eliminieren */
  background: var(--bg);
  /* damit die Hintergrundfarbe immer das ganze Fenster ausfüllt */
}

</style>
