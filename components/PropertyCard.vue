<template>
  <v-card flat class="info mt-2" style="border: 1px solid;" v-bind="$attrs" v-on="$listeners">
    <v-system-bar color="info white--text">
      {{ label }}
    </v-system-bar>
    <v-card-text class="property-text-bg pa-2" style="font-size: 0.8rem; line-height: 1.1rem;">
      <span v-html="value" />
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'PropertyCard',
  inheritAttrs: false,
  props: {
    label: { type: String, required: true },
    value: { type: String, required: true }
  }
})
</script>
