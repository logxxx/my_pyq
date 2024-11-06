<template>
  <div>
    <RouterLink to="/">Go to Home</RouterLink>
    <RouterLink to="/mgr">Go to Mgr</RouterLink>
  </div>
</template>

<script>
// @ is an alias to /src
import Index from '@/components/Index.vue'

export default {
  name: 'HomeView',
  components: {
    Index
  }
}
</script>
