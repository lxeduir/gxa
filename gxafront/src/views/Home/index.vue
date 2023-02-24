<template>
    <div id="Home">
        <Lunbotu class="lunbotu"></Lunbotu>
        <Tongyi class="tongyi"></Tongyi>
        <Taskbar class="taskbar"></Taskbar>
        <special class="special"></special>
    </div>
</template>

<script>
import Lunbotu from '../../components/Lunbotu'
import Tongyi from '../../components/Tongyi'
import Taskbar from '../../components/Taskbar'
import special from '../../components/special'
export default {
    name: 'Home',
    components: {
        Lunbotu,
        Tongyi,
        Taskbar,
        special
    }

}
</script>

<style scoped>
.special {
    position: absolute;
    height: 200px;
    top: 1200px;
    left: 50%;
    transform: translate(-50%);
    width: 1230px;
}

#Home {
    position: relative;
    width: 100vw;
    height: 1500px;
}

.Lunbotu {
    width: 100%;
    background-color: aquamarine;
}

.tongyi {
    position: absolute;
    top: 20px;
    z-index: 99;
}

.taskbar {
    /* border: 1px solid red; */
    position: relative;
    top: 450px;
    z-index: 99;
}
</style>