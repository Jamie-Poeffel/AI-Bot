<script setup>
import { ref, watch, onMounted, nextTick } from 'vue'

const userInput = ref('')
const botResponse = ref('')
const textArea = ref(null)

// Theme-Logik wie in LoginView
const theme = ref('light')

function applyTheme(t) {
  document.documentElement.setAttribute('data-theme', t)
  localStorage.setItem('theme', t)
}

onMounted(() => {
  const saved = localStorage.getItem('theme')
  if (saved) theme.value = saved
  else if (window.matchMedia('(prefers-color-scheme: dark)').matches) theme.value = 'dark'
  applyTheme(theme.value)
})


function autoResize() {
  if (textArea.value) {
    console.log('autoResize aufgerufen')
    textArea.value.style.height = 'auto'
    textArea.value.style.height = textArea.value.scrollHeight + 'px'
    console.log('neue Höhe:', textArea.value.style.height)
  }
}

watch(theme, (t) => applyTheme(t))

async function sendMessage() {
  if (!userInput.value.trim()) return
  botResponse.value = ""

  const prompt = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/prompt`, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ Text: userInput.value.trim() }),
    credentials: 'include'
  });

  const data = await prompt.json()
  botResponse.value = data?.candidates[0]?.content?.parts[0]?.text
  await nextTick()
  autoResize()
}

function formattedText() {
  let text = botResponse.value;

  text = text.replace(/\r/g, '');

  text = text.replace(/\t/g, '&nbsp;&nbsp;&nbsp;&nbsp;');

  text = text.replace(/^\* (.+)$/gm, '<li>$1</li>');

  if (text.includes('<li>')) {
    text = '<ul>' + text + '</ul>';
  }

  text = text.replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>');

  text = text.replace(/(^|[^*])\*([^*]+)\*([^*]|$)/g, '$1<em>$2</em>$3');

  text = text.replace(/\n/g, '<br>');

  return text;
}


</script>

<template>
  <div class="home-page">
    <!-- Theme-Auswahl oben rechts -->
    <div class="theme-selector">
      <select v-model="theme">
        <option value="light">Hell</option>
        <option value="dark">Dunkel</option>
      </select>
    </div>
    <div class="home-card">
      <h1>AI Bot</h1>
      <div class="chat-area">
        <div v-html="formattedText()" class="formatted-output"></div>
      </div>
      <form class="input-area" @submit.prevent="sendMessage">
        <div class="input-wrapper">
          <textarea v-model="userInput" class="input-box"
            placeholder="Stellen Sie eine Frage oder geben Sie einen Befehl ein..." rows="3"></textarea>
          <button type="submit" class="send-btn">Senden</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
.theme-selector {
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 10;
}

.theme-selector select {
  padding: .4rem .75rem;
  border-radius: .5rem;
  border: 1px solid var(--input-border, #374151);
  background: var(--input-bg, #232b3a);
  color: var(--text, #e2e8f0);
  font-size: .875rem;
}

.formatted-output {
  padding: 2rem;
  white-space: normal;
}

.unstyled-textarea {
  all: unset;
  width: 100%;
  overflow: hidden;
  resize: none;
  white-space: pre-wrap;
  font-family: inherit;
  font-size: inherit;
}

.home-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg, #181c24);
  font-family: system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif;
  color: var(--text, #e2e8f0);
  padding: 2rem;
}

.home-card {
  width: 100%;
  max-width: 80%;
  padding: 2rem;
  border-radius: 1.5rem;
  background: var(--card-bg, rgba(30, 41, 59, 0.95));
  box-shadow: 0 20px 40px rgba(0, 0, 0, .2);
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.home-card h1 {
  margin: 0 0 .5rem;
  text-align: center;
  font-size: 2rem;
  font-weight: 700;
}

.chat-area {
  min-height: 60px;
  background: var(--input-bg, #232b3a);
  border-radius: .75rem;
  padding: 1rem;
  margin-bottom: 1rem;
  color: var(--text, #e2e8f0);
  font-size: 1rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, .08);
}

.bot-message {
  white-space: pre-wrap;
  color: var(--text, #e2e8f0);
}

.input-area {
  margin-top: auto;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: flex-end;
}

.input-box {
  flex: 1;
  padding: 1rem;
  border: 2px solid var(--input-border, #374151);
  border-radius: .75rem;
  background: var(--input-bg, #232b3a);
  color: var(--text, #e2e8f0);
  font-size: 1rem;
  resize: vertical;
  transition: border-color .15s, box-shadow .15s;
  min-height: 48px;
  max-height: 120px;
  overflow-y: auto;
  margin-right: .5rem;
}

.input-box:focus {
  outline: none;
  border-color: var(--border-focus, #6366f1);
  box-shadow: 0 0 0 2px var(--border-focus, #6366f1);
}

.send-btn {
  padding: .75rem 1.25rem;
  border: none;
  border-radius: .75rem;
  background: var(--primary, #6366f1);
  color: #fff;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background .15s, opacity .15s;
  align-self: flex-end;
  margin-bottom: 4px;
}

.send-btn:hover {
  background: var(--primary-hov, #4338ca);
}
</style>