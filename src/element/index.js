import { Carousel, CarouselItem, Checkbox, CheckboxGroup, Input, Menu, MenuItem, Select,Option, Radio, DatePicker} from 'element-ui'
const element = {
    install: function (Vue) {
        Vue.use(Input)
        Vue.use(Menu)
        Vue.use(MenuItem)
        Vue.use(Carousel)
        Vue.use(CarouselItem)
        Vue.use(CheckboxGroup)
        Vue.use(Select)
        Vue.use(Checkbox)
        Vue.use(Option)
        Vue.use(Radio)
        Vue.use(DatePicker)
    }
}
export default element