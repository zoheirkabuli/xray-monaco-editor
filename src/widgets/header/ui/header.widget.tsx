import { ActionIcon, Group, Title } from '@mantine/core'
import { PiGithubLogo } from 'react-icons/pi'
import { StickyHeader } from '@/shared/ui/sticky-header'

import classes from './header.module.css'

export function HeaderWidget() {
    return (
        <StickyHeader className={classes.root} px="md">
            <Group h="100%" justify="space-between">
                <Title order={3}>Xray Config Validator</Title>

                <ActionIcon
                    component="a"
                    href="https://github.com/mmmray/xray-online"
                    size="lg"
                    target="_blank"
                    variant="subtle"
                >
                    <PiGithubLogo size={24} />
                </ActionIcon>
            </Group>
        </StickyHeader>
    )
}
