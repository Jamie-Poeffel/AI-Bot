<template>
    <div class="gpt-input-bar">
        <textarea v-model="userInput" class="gpt-textarea" placeholder="Stelle irgendeine Frage" rows="1"
            @input="autoResize" @keydown.enter.exact.prevent="sendMessage" :disabled="loading"></textarea>

        <div class="gpt-input-right">
            <button @click="sendMessage" class="icon-btn voice-btn" aria-label="Voice Input" :disabled="loading">
                <template v-if="loading">
                    <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor"
                        xmlns="http://www.w3.org/2000/svg" class="icon-lg">
                        <path
                            d="M4.5 5.75C4.5 5.05964 5.05964 4.5 5.75 4.5H14.25C14.9404 4.5 15.5 5.05964 15.5 5.75V14.25C15.5 14.9404 14.9404 15.5 14.25 15.5H5.75C5.05964 15.5 4.5 14.9404 4.5 14.25V5.75Z">
                        </path>
                    </svg> </template>
                <template v-else>
                    <ArrowUp class="lucide-icon" />
                </template>
            </button>
        </div>
    </div>
</template>

<script setup>
import { ref, watch, defineProps } from 'vue'
import {
    ArrowUp,
} from 'lucide-vue-next'

const emit = defineEmits(['send'])

const props = defineProps({
    loading: {
        type: Boolean,
        default: false
    }
})

const userInput = ref('')

const autoResize = (e) => {
    const el = e.target
    el.style.height = 'auto'
    el.style.height = `${el.scrollHeight}px`
}

const sendMessage = () => {
    if (props.loading || !userInput.value.trim()) return
    props.loading = true

    emit('send', userInput.value)
    userInput.value = ''
}

</script>

<style scoped>
.gpt-input-bar {
    display: flex;
    align-items: center;
    background: rgba(255, 255, 255, 0.12);
    border: 1px solid rgba(255, 255, 255, 0.2);
    border-radius: 999px;
    padding: 0.75rem 1rem;
    gap: 1rem;
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    width: 100%;
    margin: 2rem 0;
    color: var(--text);
}

.gpt-input-left,
.gpt-input-right {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

.gpt-textarea {
    flex: 1;
    resize: none;
    border: none;
    outline: none;
    font-size: 1rem;
    background: transparent;
    min-height: 24px;
    line-height: 1.4;
    padding: 0;
    overflow: hidden;
    color: var(--text);
}

.gpt-textarea:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.icon-btn {
    background: none;
    border: none;
    padding: 0.4rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    border-radius: 0.5rem;
    transition: background 0.2s ease;
    color: var(--text);
}

.icon-btn:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.icon-btn:hover:enabled {
    background: rgba(255, 255, 255, 0.1);
}

.voice-btn {
    background: rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(0, 0, 0, 0.3);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border-radius: 50%;
    width: 32px;
    height: 32px;
    justify-content: center;
}

.lucide-icon {
    width: 20px;
    height: 20px;
    stroke: var(--text);
}

.loading-block {
    display: block;
    width: 18px;
    height: 18px;
    background: var(--text);
    border-radius: 4px;
    opacity: 0.5;
    animation: loading-blink 1s infinite alternate;
}

@keyframes loading-blink {
    0% {
        opacity: 0.5;
    }

    100% {
        opacity: 1;
    }
}
</style>
