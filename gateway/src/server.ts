import express from 'express'
import swaggerUi from 'swagger-ui-express'
import swaggerSpec from './swaggerOptions'
import { setupRoutes } from './routes'

const app = express()
const port = 3000

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec))

app.use(express.json())

setupRoutes(app)

app.listen(port, () => {
    console.log(`API Gateway running at http://localhost:${port}`)
})
