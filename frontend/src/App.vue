<script setup>
import { onMounted, ref } from 'vue';

import ButtonClose from '@/components/ButtonClose.vue'
import SelectionRegion from '@/components/SelectionRegion.vue'
import FormSettings from '@/components/FormSettings.vue'

import {Button} from "@/components/ui/button"
import {TakeScreenshot} from '../wailsjs/go/main/App'

const submitting = ref(false)

const start = ref(null)
const end = ref(null)

const settingsOpen = ref(false)

const handleUpdateRegion = (updatedStart, updatedEnd) => {
  start.value = updatedStart
  end.value = updatedEnd
}

const captureScreen = async () => {
  submitting.value = true

  // Hacky race condition handling, lets not do this long term
  await new Promise(r => setTimeout(r, 500))

  await TakeScreenshot(
    start.value.x, 
    start.value.y, 
    end.value.x, 
    end.value.y
  )

  submitting.value = false
}

</script>
<template>
  <div class="relative w-full h-full">
    <SelectionRegion :submitting="submitting" @update-region="handleUpdateRegion"/>

    <div class="absolute top-0 left-[50%] -translate-x-1/2 flex items-center justify-center mt-8 text-white">
      <div class="relative border bg-background rounded-2xl px-6 py-4">
        <ButtonClose class="absolute -top-3 -right-3"/>
  
        <!-- These different capture types are not yet implemented -->
        <!-- Only "selection" exists at this point -->
        <div class="grid grid-cols-3 gap-2">
          <button class="flex flex-col items-center justify-center px-4 py-2 rounded-lg bg-gray-100/20">
            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-device-desktop w-14" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
              <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
              <path d="M8 5v10a1 1 0 0 0 1 1h10" />
              <path d="M5 8h10a1 1 0 0 1 1 1v10" />
            </svg>
            Selection
          </button>
  
          <button class="flex flex-col items-center justify-center px-4 py-2 rounded-lg hover:bg-gray-100/10">
            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-device-desktop w-14" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
              <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
              <path d="M3 5a1 1 0 0 1 1 -1h16a1 1 0 0 1 1 1v10a1 1 0 0 1 -1 1h-16a1 1 0 0 1 -1 -1v-10z" />
              <path d="M7 20h10" />
              <path d="M9 16v4" />
              <path d="M15 16v4" />
            </svg>
            Screen
            <span class="text-xs text-gray-300 mt-2">Coming soon</span>
          </button>
  
          <button class="flex flex-col items-center justify-center px-4 py-2 rounded-lg hover:bg-gray-100/10">
            <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-device-desktop w-14" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
              <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
              <path d="M3 5m0 2a2 2 0 0 1 2 -2h14a2 2 0 0 1 2 2v10a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2z" />
              <path d="M6 8h.01" />
              <path d="M9 8h.01" />
            </svg>
            Window
            <span class="text-xs text-gray-300 mt-2">Coming soon</span>
          </button>
        </div>
        <div class="grid grid-cols-3 text-center mt-4">
          <div class="flex justify-center items-center flex-col">
            <div class="flex items-center border w-min p-1 rounded-lg space-x-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-polaroid w-6 bg-gray-300 stroke-black rounded" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                <path d="M4 4m0 2a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2z" />
                <path d="M4 16l16 0" />
                <path d="M4 12l3 -3c.928 -.893 2.072 -.893 3 0l4 4" />
                <path d="M13 12l2 -2c.928 -.893 2.072 -.893 3 0l2 2" />
                <path d="M14 7l.01 0" />
              </svg>
              <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-gif w-6" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                <path d="M8 8h-2a2 2 0 0 0 -2 2v4a2 2 0 0 0 2 2h2v-4h-1" />
                <path d="M12 8v8" />
                <path d="M16 12h3" />
                <path d="M20 8h-4v8" />
              </svg>
            </div>
            <span class="text-xs text-gray-300 mt-2">Coming soon</span>
          </div>
  
          <div class="flex justify-center">
            <button 
              class="w-12 h-12 rounded-full outline outline-offset-2 flex items-center justify-center" 
              :class="submitting ? 'bg-background' : 'bg-white hover:bg-gray-300'"
              :disabled="submitting"
              @click.prevent="captureScreen">
              <svg v-if="submitting" xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-loader-2 animate-spin w-10" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                <path d="M12 3a9 9 0 1 0 9 9" />
              </svg>
            </button>
          </div>
  
          <div class="flex justify-center items-center">
            <Button 
              variant="ghost"
              :class="settingsOpen ? 'bg-accent' : ''"
              @click.prevent="settingsOpen = !settingsOpen">
              <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-tool w-5" viewBox="0 0 24 24" stroke-width="1.5" stroke="#ffffff" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                <path d="M7 10h3v-3l-3.5 -3.5a6 6 0 0 1 8 8l6 6a2 2 0 0 1 -3 3l-6 -6a6 6 0 0 1 -8 -8l3.5 3.5" />
              </svg>
            </Button>
          </div>
        </div>

        <div v-if="settingsOpen" class="bg-gray-100/10 mt-4 rounded-lg py-4 px-2">
          <h3 class="font-semibold">Settings</h3>
          <FormSettings/>
        </div>
      </div>
    </div>
  </div>
</template>
