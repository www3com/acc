import style from './style.less'

interface Dialog {
    visible?: boolean
    children: any
}

export default ({visible, children}: Dialog) => {
    return (<div className={style.fullDialog} style={{display: visible && visible == true ? 'block' : 'none'}}>
        {children}
    </div>);
};
