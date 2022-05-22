import BCCHeader from './BCCHeader'
import { BCCFooter } from './BCCFooter'

const layoutStyle = {
    display: 'flex',
    flexDirection: 'column',
    height: '100%',
    width: '100%',
}

const contentStyle = {
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
}

const BCCLayout = ({ children }) => {
    return (
        <div className="layout" style={layoutStyle}>
            <BCCHeader />
            <div className="content" style={contentStyle}>
                    {children}
            </div>
            <BCCFooter />
        </div>
    )
}

export default BCCLayout
