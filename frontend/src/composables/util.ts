// 异步调用状态控制 防止重复点击的封装
// @ts-ignore
import { EventsOn } from '@runtime'

export function useAsync(fn: any) {
  const pending = ref(false)

  const func = async (...args: any[]) => {
    if (pending.value !== false) return;
    pending.value = true

    await fn(...args)

    pending.value = false
  }

  return reactive({ pending, func })
}
