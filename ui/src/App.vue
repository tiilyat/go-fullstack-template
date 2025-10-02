<script setup lang="ts">
import { ref, onMounted } from 'vue'

const backendStatus = ref('')

onMounted(async () => {
  const response = await fetch('/api/ping')
  const data = await response.json()
  backendStatus.value = data.message
})
</script>

<template>
  <h1>You did it</h1>
  <p>Backend status: {{ backendStatus }}</p>
  <p>
    Visit <a href="https://vuejs.org/" target="_blank" rel="noopener">vuejs.org</a> to read the
    documentation
  </p>
</template>
