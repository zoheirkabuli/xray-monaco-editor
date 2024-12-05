import { PiCheckSquareOffset } from 'react-icons/pi'
import { Button, Group } from '@mantine/core'

import { Props } from './interfaces'

export function ConfigEditorActionsFeature(props: Props) {
    const { editorRef } = props

    const formatDocument = () => {
        if (!editorRef.current) return
        if (typeof editorRef.current !== 'object') return
        if (!('getAction' in editorRef.current)) return
        if (typeof editorRef.current.getAction !== 'function') return

        editorRef.current.getAction('editor.action.formatDocument').run()
    }

    return (
        <Group>
            <Button
                leftSection={<PiCheckSquareOffset size={16} />}
                mb="md"
                onClick={formatDocument}
            >
                Format
            </Button>
        </Group>
    )
}
