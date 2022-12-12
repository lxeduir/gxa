<template>
    <div class="swiper-container">
        <div class="zhuanqu">
            <span>特色专区</span>
        </div>
        <div class="swiper-wrapper">
            <div class="swiper-slide"><img src="@/assets/1.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/2.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/3.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/4.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/5.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/6.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/7.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/8.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/9.png" alt="" height="160" width="277"></div>
            <div class="swiper-slide"><img src="@/assets/10.png" alt="" height="160" width="277"></div>
        </div>
        <!-- 分页器 -->
        <div class="swiper-pagination"></div>
        <!-- 导航按钮 -->
        <div class="swiper-button-prev"></div>
        <div class="swiper-button-next"></div>
    </div>
</template>
<script>
import 'swiper/css/swiper.css'
import Swiper from 'swiper'
export default {
    name: 'carrousel-items',
    data() {
        return {
            banner: []
        }
    },
    mounted() {
        new Swiper('.swiper-container', {
            loop: true,

            pagination: {
                el: '.swiper-pagination',
            },
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev',
            },
            slidesPerView: 5,
            loopedSlides: 5,
            autoplay: {
                delay: 2000,
                stopOnLastSlide: false,
                disableOnInteraction: true,
            }
        })
    }
}
</script>
<style>
.zhuanqu {
    width: 200px;
    height: 50px;
    position: relative;
}
.zhuanqu span {
    font-size: 30px;
    position: absolute;
    left: 10px;
    font-weight: 900;
}
</style>