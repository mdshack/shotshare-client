<script setup>
import {Button} from "@/components/ui/button"
import {Input} from "@/components/ui/input"
import {Label} from "@/components/ui/label"
</script>
<template>
    <form class="pt-4 space-y-2">
        <div>
            <Label>ShotShare Url</Label>
            <Input/>
        </div>
        <div>
            <Label>API Key</Label>
            <Input/>
        </div>
        <div class="pt-2">
            <Button class="w-full">Update</Button>
        </div>
    </form>
</template>