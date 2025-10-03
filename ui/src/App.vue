<script setup lang="ts">
import { ref } from 'vue'
import {Button} from "@/components/ui/button";

const backendStatus = ref('')

function checkBackendStatus() {
  fetch('/api/health')
    .then(response => response.json())
    .then(data => {
      backendStatus.value = data.message === 'ok' ? 'Healthy' : 'Unhealthy'
    })
    .catch(() => {
      backendStatus.value = 'Error'
    })
}
</script>

<template>
  <Button variant="default" @click="checkBackendStatus">Check backend status</Button>
  <p v-if="Boolean(backendStatus)">Backend status: {{ backendStatus }}</p>
</template>
