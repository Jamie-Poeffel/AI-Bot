<!-- LoginView.vue -->
<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'

/* ---------- Form-State ---------- */
const email    = ref('')
const password = ref('')
const errorMsg = ref('')
const loading  = ref(false)
const router   = useRouter()

/* ---------- Theme-State ---------- */
const theme = ref<'light' | 'dark'>('light')

/* lies gespeichertes Theme oder System-Präferenz */
onMounted(() => {
  const saved = localStorage.getItem('theme') as 'light' | 'dark' | null
  if (saved) {
    theme.value = saved
  } else if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    theme.value = 'dark'
  }
  applyTheme(theme.value)
})

/* wende Theme auf <html data-theme=""> an und speichere es */
function applyTheme(t: 'light' | 'dark') {
  document.documentElement.setAttribute('data-theme', t)
  localStorage.setItem('theme', t)
}

/* auf Änderung reagieren */
watch(theme, (t) => applyTheme(t))

/* ---------- Login ---------- */
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
    router.push('/')                       // Zielroute anpassen
  } catch (e: any) {
    errorMsg.value = e.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <!-- ---------- Theme-Dropdown ---------- -->
  <div class="theme-selector">
    <select v-model="theme">
      <option value="light">Hell</option>
      <option value="dark">Dunkel</option>
    </select>
  </div>

  <!-- ---------- Login ---------- -->
  <div class="login-page">
    <div class="login-card">
      <h1>Login</h1>

      <form @submit.prevent="handleLogin">
        <label for="email">E-Mail</label>
        <input
          id="email"
          v-model="email"
          type="email"
          autocomplete="username"
          placeholder="name@example.com"
          required
        />

        <label for="password">Passwort</label>
        <input
          id="password"
          v-model="password"
          type="password"
          autocomplete="current-password"
          placeholder="••••••••"
          required
        />

        <button :disabled="loading">
          {{ loading ? 'Anmelden …' : 'Anmelden' }}
        </button>

        <p v-if="errorMsg" class="error">{{ errorMsg }}</p>
      </form>
    </div>
  </div>
</template>

<style scoped>
/* ---------- Grund­farben (Light) ---------- */
:root {
  --bg            : #f0f4ff;
  --card-bg       : rgba(255, 255, 255, 0.95);
  --input-bg      : #ffffff;
  --input-border  : #d1d5db;
  --border-focus  : #4f46e5;
  --text          : #0f172a;
  --muted         : #64748b;
  --primary       : #4f46e5;
  --primary-hov   : #4338ca;
  --error         : #ef4444;
}

/* ---------- Dark-Overrides ---------- */
[data-theme='dark'] {
  --bg            : #0f172a;
  --card-bg       : rgba(30, 41, 59, 0.8);
  --input-bg      : rgba(51, 65, 85, 0.6);
  --input-border  : #64748b;
  --text          : #e2e8f0;
  --muted         : #94a3b8;
}

/* ---------- Theme-Dropdown ---------- */
.theme-selector {
  position : fixed;
  top      : 1rem;
  right    : 1rem;
  z-index  : 10;
}

.theme-selector select {
  padding       : 0.4rem 0.75rem;
  border-radius : 0.5rem;
  border        : 1px solid var(--input-border);
  background    : var(--input-bg);
  color         : var(--text);
  font-size     : 0.875rem;
}

/* ---------- Login-Layout ---------- */
.login-page {
  min-height : 100vh;
  display    : flex;
  align-items: center;
  justify-content: center;
  background : var(--bg);
  font-family: system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif;
  color      : var(--text);
  padding    : 2rem;
}

.login-card {
  width          : 100%;
  max-width      : 24rem;
  background     : var(--card-bg);
  backdrop-filter: blur(8px);
  padding        : 2rem;
  border-radius  : 1.5rem;
  box-shadow     : 0 20px 40px rgba(0, 0, 0, 0.2);
}

.login-card h1 {
  margin      : 0 0 1.5rem;
  text-align  : center;
  font-size   : 2rem;
  font-weight : 700;
}

.login-card form {
  display       : flex;
  flex-direction: column;
  gap           : 1.25rem;
}

.login-card label {
  font-size  : 0.875rem;
  font-weight: 500;
  color      : var(--muted);
}

.login-card input {
  width        : 100%;
  padding      : 0.5rem 1rem;
  border       : 2px solid var(--input-border);
  border-radius: 0.75rem;
  background   : var(--input-bg);
  color        : var(--text);
  font-size    : 1rem;
  transition   : border-color 0.15s, box-shadow 0.15s;
}

.login-card input::placeholder { color: var(--muted); }

.login-card input:focus {
  outline     : none;
  border-color: var(--border-focus);
  box-shadow  : 0 0 0 2px var(--border-focus);
}

.login-card button {
  padding      : 0.75rem;
  border       : none;
  border-radius: 0.75rem;
  background   : var(--primary);
  color        : #fff;
  font-size    : 1rem;
  font-weight  : 600;
  cursor       : pointer;
  transition   : background 0.15s, opacity 0.15s;
}

.login-card button:hover    { background: var(--primary-hov); }
.login-card button:disabled { opacity   : 0.5; cursor: not-allowed; }

.error {
  text-align : center;
  font-size  : 0.875rem;
  color      : var(--error);
}
</style>
