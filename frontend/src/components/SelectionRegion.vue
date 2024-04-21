<script setup>
import {ref, onMounted} from 'vue'

const emit = defineEmits([
    'updateRegion'
])

defineProps({
  submitting: Boolean
})

const start = ref({
  x: null,
  y: null,
})

const end = ref({
  x: null,
  y: null,
})

const canvas = ref(null)

const setCanvasSize = () => {
  canvas.value.width = window.innerWidth
  canvas.value.height = window.innerHeight
}

window.addEventListener('resize', setCanvasSize, false)
onMounted(setCanvasSize)

const drawingSelection = ref(false)

const clearCanvas = (ctx = null) => {
  if(!ctx) {
    ctx = canvas.value.getContext("2d")
  }
  ctx.clearRect(0,0,canvas.value.width, canvas.value.height)
}

const onPointerDown = (event) => {
  drawingSelection.value = true
  start.value.x = event.clientX
  start.value.y = event.clientY
}

const onPointerUp = (event) => {
  drawingSelection.value = false

  end.value.x = event.clientX
  end.value.y = event.clientY

  emit('updateRegion', start.value, end.value)
}

const onPointerCancel = (event) => {
  console.log("cancel")
  console.log(event)
}

const onPointerMove = (event) => {
  if(!drawingSelection.value) return

  const ctx = canvas.value.getContext("2d")

  clearCanvas(ctx)

  ctx.lineWidth = 3
  ctx.strokeStyle = "#fff"
  ctx.fillStyle = "rgba(255,255,255,50%)";

  ctx.beginPath()
  ctx.fillRect(
    start.value.x, 
    start.value.y, 
    event.clientX - start.value.x , 
    event.clientY - start.value.y, 
  )
  ctx.stroke()
}
</script>

<template>
    <canvas 
        ref="canvas"
        class="w-full h-full bg-background/20"
        :class="submitting ? 'hidden' : ''"
        @pointerdown.self="onPointerDown"
        @pointerup.self="onPointerUp"
        @pointercancel.self="onPointerCancel"
        @pointerout.self="onPointerCancel"
        @pointermove="onPointerMove">
    </canvas>
</template>