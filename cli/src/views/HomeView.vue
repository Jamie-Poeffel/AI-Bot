<script setup>
import { ref, watch, onMounted, nextTick } from 'vue'
import Chatimput from '../components/chatimput.vue'

const textArea = ref(null)
const messages = ref([])
const load = ref(false)
const chatArea = ref(null)

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

async function sendMessage(val) {
  load.value = true
  messages.value.push({
    role: 'user',
    text: val
  });
  scrollToBottom()

  const prompt = await fetch(`${import.meta.env.VITE_APP_BASE_URL}/prompt`, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ Text: val }),
    credentials: 'include'
  });

  const data = await prompt.json()
  const newVal = data?.candidates[0]?.content?.parts[0]?.text
  messages.value.push({
    role: 'bot',
    text: newVal
  });
  await nextTick()
  autoResize()
  scrollToBottom()
  load.value = false
}

function scrollToBottom() {
  if (chatArea.value) {
    chatArea.value.scrollTop = chatArea.value.scrollHeight
  }
}

function formattedText(input) {
  if (!input || !input.trim()) return '';
  let text = input;

  text = text.replace(/\r/g, '');

  text = text.replace(/\t/g, '&nbsp;&nbsp;&nbsp;&nbsp;');

  // Handle multiline code blocks (```...```)
  text = text.replace(/```([\s\S]*?)```/g, function (match, code) {
    return `<pre><code style="{ max-width: 90%; }">${code.replace(/</g, '&lt;').replace(/>/g, '&gt;')}</code></pre>`;
  });

  // Handle inline code (`...`)
  text = text.replace(/`([^`\n]+?)`/g, function (match, code) {
    return `<code>${code.replace(/</g, '&lt;').replace(/>/g, '&gt;')}</code>`;
  });

  // Handle list items
  text = text.replace(/^\* (.+)$/gm, '<li style="{ transform: translate(10px); }">$1</li>');
  if (text.includes('<li>')) {
    text = '<ul>' + text + '</ul>';
  }

  // Bold (**text**)
  text = text.replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>');

  // Italic (*text*) - but not inside code blocks
  // First, temporarily replace code blocks and inline code with placeholders
  const codeBlocks = [];
  text = text.replace(/<pre><code>[\s\S]*?<\/code><\/pre>|<code>[^<]*<\/code>/g, function (match) {
    codeBlocks.push(match);
    return `__CODE_BLOCK_${codeBlocks.length - 1}__`;
  });

  // Italic (*text*)
  text = text.replace(/(^|[^*])\*([^*]+)\*([^*]|$)/g, '$1<em>$2</em>$3');

  // Restore code blocks
  text = text.replace(/__CODE_BLOCK_(\d+)__/g, (m, i) => codeBlocks[i]);

  // Newlines
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
    <div class="chat-area" ref="chatArea">
      <div v-for="(msg, index) in messages" :key="index" class="message"
        :class="msg.role === 'user' ? 'user-message' : 'bot-message'" v-html="formattedText(msg.text)"></div>
    </div>

    <form class="input-area" @submit.prevent="sendMessage">
      <Chatimput @send="(val) => sendMessage(val)" :loading="load" />
    </form>
  </div>
</template>

<style scoped>
.home-page {
  display: flex;
  flex-direction: column;
  height: 100vh;
  /* full viewport height */
  padding: 2rem;
  background: var(--bg, #181c24);
  font-family: system-ui, -apple-system, 'Segoe UI', Roboto, sans-serif;
  color: var(--text, #e2e8f0);
  position: relative;
  overflow: hidden;
}

.theme-selector {
  position: fixed;
  top: 1rem;
  right: 1rem;
  z-index: 10;
}

.theme-selector select {
  padding: 0.4rem 0.75rem;
  border-radius: 0.5rem;
  border: 1px solid var(--input-border, #374151);
  background: var(--input-bg, #232b3a);
  color: var(--text, #e2e8f0);
  font-size: 0.875rem;
}

.chat-area {
  flex: 1 1 auto;
  /* grow and shrink */
  overflow-y: auto;
  margin-bottom: 1rem;
  max-width: 800px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding: 1rem;
  background: transparent;
  scrollbar-width: thin;
  scrollbar-color: rgba(16, 163, 127, 0.5) transparent;
  color: var(--text);
}

.chat-area::-webkit-scrollbar {
  width: 8px;
}

.message ul li {
  transform: translateX(10px) !important;
}

.chat-area::-webkit-scrollbar-thumb {
  background-color: rgba(16, 163, 127, 0.5);
  border-radius: 4px;
}

.message {
  background: rgba(255, 255, 255, 0.12);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 19px;
  padding: 0.75rem 1rem;
  color: var(--text);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  white-space: pre-wrap;
  font-size: 1rem;
  line-height: 1.4;
  max-width: 80%;
  word-break: break-word;
  width: fit-content;
}

/* Differentiate user vs bot */
.bot-message {
  align-self: flex-start;
  background: rgba(255, 255, 255, 0.12);
}

.user-message {
  align-self: flex-end;
  background: rgba(16, 163, 127, 0.3);
  border-color: rgba(16, 163, 127, 0.5);
}

.input-area {
  flex-shrink: 0;
  max-width: 800px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  display: flex;
  padding-top: 0.5rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}
</style>