import '@mantine/carousel/styles.layer.css'
import '@mantine/charts/styles.layer.css'
import '@mantine/code-highlight/styles.layer.css'
import '@mantine/core/styles.layer.css'
import '@mantine/dates/styles.layer.css'
import '@mantine/dropzone/styles.layer.css'
import '@mantine/notifications/styles.layer.css'
import '@mantine/nprogress/styles.layer.css'

import './global.css'

import { NavigationProgress } from '@mantine/nprogress'
import { MantineProvider } from '@mantine/core'

import { ConfigPageConnector } from '@pages/config/ui/connectors/config.page.connector'

export function App() {
    return (
        <MantineProvider defaultColorScheme="dark">
            <NavigationProgress />
            <ConfigPageConnector />
        </MantineProvider>
    )
}
