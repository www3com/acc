import {Button} from "antd";
import {useEffect, useState} from "react";

export default (props: { onClick?: any, children: any }) => {
  const [children, setChildren] = useState<string>(props.children);
  const [disabled, setDisabled] = useState<boolean>(false)

  let timer: NodeJS.Timer, second = 60;

  useEffect(() => {
    return () => {
      clear()
    }
  }, [])

  const clear = () => {
    window.clearInterval(timer)
    setChildren(props.children)
    setDisabled(false)
  }

  const run = () => {
    setChildren(--second + '秒后重新获取')
    if (second == 0) {
      clear()
    }
  }

  const onClick = () => {
    if (props.onClick) {
      if (!props.onClick()) {
        return
      }
    }

    setDisabled(true)
    timer = setInterval(run, 1000)
  }

  return (<Button disabled={disabled} onClick={onClick}>{children}</Button>)
}
