<script setup lang="ts">
const props = defineProps<{ imageUrl: string | null; loading: boolean }>()

function downloadQRCode() {
  if (!props.imageUrl) return
  const link = document.createElement('a')
  link.href = props.imageUrl
  link.download = 'qrcode.png'
  link.click()
}
</script>

<template>
  <div class="flex flex-col justify-center items-center h-full gap-2">
    <div v-if="props.loading" class="text-gray-500 italic">Generating QR-Code...</div>
    <img
      v-else-if="props.imageUrl"
      :src="props.imageUrl"
      alt="QR Code"
      class="max-w-full max-h-[80vh]"
    />
    <div v-else class="text-gray-400 italic">No QR-Code generated yet</div>

    <button
      v-if="props.imageUrl"
      class="bg-teal-600 text-white rounded p-2 hover:bg-teal-700 mt-2"
      @click="downloadQRCode"
    >
      QR-Code herunterladen
    </button>
  </div>
</template>
