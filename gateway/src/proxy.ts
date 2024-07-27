import axios from 'axios'
import { Request, Response } from 'express'

export const proxyRequest = async (
    serviceUrl: string,
    prefix: string,
    req: Request,
    res: Response,
) => {
    const targetPath = req.originalUrl.replace(new RegExp(`^/${prefix}`), '')

    console.log(`Forwarding request to: ${serviceUrl}${targetPath}`)

    try {
        const response = await axios({
            method: req.method,
            url: `${serviceUrl}${targetPath}`,
            data: req.body,
        })

        res.set('Cache-Control', 'no-store')
        res.set('Pragma', 'no-cache')
        res.set('Expires', '0')

        console.log(`Response from backend: ${response.status}`)
        res.status(response.status).send(response.data)
    } catch (error: any) {
        const formattedError = formatErrorResponse(error)
        res.status(error.response?.status || 500).json(formattedError)
    }
}

const formatErrorResponse = (error: any) => {
    if (error.response && error.response.data) {
        return {
            data: {
                error: error.response.data.error || error.message,
            },
        }
    }
    return {
        data: {
            error: error.message,
        },
    }
}
