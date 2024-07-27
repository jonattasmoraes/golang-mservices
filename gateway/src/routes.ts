import { Application } from 'express'
import { proxyRequest } from './proxy'
import dotenv from 'dotenv'

dotenv.config()

export const setupRoutes = (app: Application) => {
    /**
     * @openapi
     * /user/api/users/:
     *   get:
     *     tags:
     *       - User
     *     summary: Request all users
     *     description: This route proxies requests to the user service with pagination.
     *     parameters:
     *       - in: query
     *         name: page
     *         schema:
     *           type: integer
     *         description: Page number for pagination
     *         required: false
     *     responses:
     *       200:
     *         description: Successfully proxied request to user service
     *       404:
     *         description: Not found
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       500:
     *         description: Internal server error
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     */
    app.use('/user', (req, res) =>
        proxyRequest(process.env.USER_SERVICE_URL as string, 'user', req, res),
    )

    /**
     * @openapi
     * /user/api/users:
     *   post:
     *     tags:
     *       - User
     *     summary: Create a new user
     *     description: This route creates a new user with the provided details.
     *     requestBody:
     *       required: true
     *       content:
     *         application/json:
     *           schema:
     *             type: object
     *             properties:
     *               first_name:
     *                 type: string
     *                 example: John
     *                 description: The first name of the user.
     *               last_name:
     *                 type: string
     *                 example: Doe
     *                 description: The last name of the user.
     *               email:
     *                 type: string
     *                 example: john.doe@example.com
     *                 description: The email address of the user.
     *               password:
     *                 type: string
     *                 example: my_secure_password
     *                 description: The password for the user account.
     *             required:
     *               - first_name
     *               - last_name
     *               - email
     *               - password
     *     responses:
     *       201:
     *         description: User successfully created
     *         content:
     *           application/json:
     *             schema:
     *               type: object
     *               properties:
     *                 data:
     *                   type: object
     *                   properties:
     *                     id:
     *                       type: string
     *                       description: The unique identifier for the created user.
     *                       example: 01J3RP0N8AHRK01C2ZD8TW6451
     *                     first_name:
     *                       type: string
     *                       description: The first name of the user.
     *                       example: John
     *                     last_name:
     *                       type: string
     *                       description: The last name of the user.
     *                       example: Doe
     *                     email:
     *                       type: string
     *                       description: The email address of the user.
     *                       example: john.doe@example.com
     *                     role:
     *                       type: string
     *                       description: The role assigned to the user.
     *                       example: user
     *                     create_at:
     *                       type: string
     *                       format: date-time
     *                       description: The timestamp when the user was created.
     *                       example: '2024-07-26T23:29:00Z'
     *                     update_at:
     *                       type: string
     *                       format: date-time
     *                       description: The timestamp when the user was last updated.
     *                       example: '2024-07-26T23:29:00Z'
     *                 message:
     *                   type: string
     *                   description: Success message indicating the operation result.
     *                   example: 'operation from handler: create user, successful.'
     *       400:
     *         description: Bad request, invalid input
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       500:
     *         description: Internal server error
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     * components:
     *   schemas:
     *     ErrorResponse:
     *       type: object
     *       properties:
     *         code:
     *           type: integer
     *           description: Error code.
     *           example: 500
     *         message:
     *           type: string
     *           description: Error message.
     *           example: Internal server error
     */
    app.use('/user/api/users', (req, res) =>
        proxyRequest(process.env.USER_SERVICE_URL as string, 'user', req, res),
    )

    /**
     * @openapi
     * /product/api/products/:
     *   get:
     *     tags:
     *       - Product
     *     summary: Request a list of products
     *     description: This route proxies requests to the product service with pagination.
     *     parameters:
     *       - in: query
     *         name: page
     *         schema:
     *           type: integer
     *         description: Page number for pagination
     *         required: false
     *     responses:
     *       200:
     *         description: Successfully proxied request to product service
     *       404:
     *         description: Not found
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       500:
     *         description: Internal server error
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     */
    app.use('/product', (req, res) =>
        proxyRequest(process.env.PRODUCT_SERVICE_URL as string, 'product', req, res),
    )

    /**
     * @openapi
     * /sale/sales:
     *   post:
     *     tags:
     *       - Sale
     *     summary: Create a new sale
     *     description: This endpoint proxies requests to the sale service to create a new sale with the provided details.
     *     requestBody:
     *       required: true
     *       content:
     *         application/json:
     *           schema:
     *             type: object
     *             properties:
     *               user_id:
     *                 type: string
     *                 description: Unique identifier for the user making the sale
     *                 example: "01J1P83E4KJ3EM3H73139AC5PQ"
     *               payment_type:
     *                 type: string
     *                 description: Payment method used for the sale
     *                 example: "credit"
     *               products:
     *                 type: array
     *                 items:
     *                   type: object
     *                   properties:
     *                     product_id:
     *                       type: string
     *                       description: Unique identifier for the product
     *                       example: "01H7D8E4K5TX7P9ZP6G3V0A1W2"
     *                     quantity:
     *                       type: integer
     *                       description: Quantity of the product sold
     *                       example: 2
     *             required:
     *               - user_id
     *               - payment_type
     *               - products
     *     responses:
     *       201:
     *         description: Successfully created the sale
     *         content:
     *           application/json:
     *             schema:
     *               type: object
     *               properties:
     *                 data:
     *                   type: object
     *                   properties:
     *                     UserID:
     *                       type: string
     *                       description: Unique identifier for the user who made the sale
     *                     SaleID:
     *                       type: string
     *                       description: Unique identifier for the created sale
     *                     PaymentType:
     *                       type: string
     *                       description: Payment method used for the sale
     *                     Date:
     *                       type: string
     *                       format: date-time
     *                       description: Date and time when the sale occurred
     *                     Products:
     *                       type: array
     *                       items:
     *                         type: object
     *                         properties:
     *                           ProductID:
     *                             type: string
     *                             description: Unique identifier for the product
     *                           Name:
     *                             type: string
     *                             description: Name of the product
     *                           Unit:
     *                             type: string
     *                             description: Unit of measurement for the product
     *                           Category:
     *                             type: string
     *                             description: Category of the product
     *                           Quantity:
     *                             type: integer
     *                             description: Quantity of the product sold
     *                           Price:
     *                             type: integer
     *                             description: Price of the product
     *             examples:
     *               application/json:
     *                 summary: Sample response
     *                 value:
     *                   data:
     *                     UserID: "01J1P83E4KJ3EM3H73139AC5PQ"
     *                     SaleID: "01J3RVWJNDABEQ1KK4VW12BXV4"
     *                     PaymentType: "credit"
     *                     Date: "2024-07-26 22:11:38"
     *                     Total: 6
     *                     Products:
     *                       - ProductID: "01H7D8E4K5TX7P9ZP6G3V0A1W2"
     *                         Name: "Leite UHT 1L"
     *                         Unit: "Litro"
     *                         Category: "Bebidas"
     *                         Quantity: 2
     *                         Price: 3
     *       400:
     *         description: Bad request due to invalid input
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       500:
     *         description: Internal server error
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     */
    app.use('/sale', (req, res) =>
        proxyRequest(process.env.SALE_SERVICE_URL as string, 'sale', req, res),
    )

    /**
     * @openapi
     * /report/api/v1/reports/{sale_id}:
     *   get:
     *     tags:
     *       - Report
     *     summary: Request a sale by ID
     *     description: This endpoint proxies requests to the report service to retrieve a report by the specified sale ID.
     *     parameters:
     *       - in: path
     *         name: sale_id
     *         schema:
     *           type: string
     *         description: Unique identifier for the sale to retrieve the report for
     *         required: true
     *     responses:
     *       200:
     *         description: Successfully retrieved the report for the given sale ID
     *         content:
     *           application/json:
     *             schema:
     *               type: object
     *               properties:
     *                 UserID:
     *                   type: string
     *                   description: Unique identifier for the user who made the sale
     *                 SaleID:
     *                   type: string
     *                   description: Unique identifier for the sale
     *                 PaymentType:
     *                   type: string
     *                   description: Payment method used for the sale
     *                 date:
     *                   type: string
     *                   format: date-time
     *                   description: Date and time when the sale occurred
     *                 Total:
     *                   type: integer
     *                   description: Total amount of the sale
     *                 Products:
     *                   type: array
     *                   items:
     *                     type: object
     *                     properties:
     *                       ProductID:
     *                         type: string
     *                         description: Unique identifier for the product
     *                       Name:
     *                         type: string
     *                         description: Name of the product
     *                       Unit:
     *                         type: string
     *                         description: Unit of measurement for the product
     *                       Category:
     *                         type: string
     *                         description: Category of the product
     *                       Quantity:
     *                         type: integer
     *                         description: Quantity of the product sold
     *                       Price:
     *                         type: integer
     *                         description: Price of the product
     *             examples:
     *               application/json:
     *                 summary: Sample response
     *                 value:
     *                   UserID: "01J1P83E4KJ3EM3H73139AC5PQ"
     *                   SaleID: "01J3RN9RAGGGM1WWVHY2P5BEDW"
     *                   PaymentType: "credit"
     *                   date: "2024-07-26T20:16:30.160995Z"
     *                   Total: 62
     *                   Products:
     *                     - ProductID: "01H7D8E4M7TX2R3ZQ8G5V2C3Y4"
     *                       Name: "Refrigerante Coca Cola Lata 350ml"
     *                       Unit: "un"
     *                       Category: "Bebidas"
     *                       Quantity: 2
     *                       Price: 8
     *       400:
     *         description: Bad request due to missing or invalid sale_id
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       404:
     *         description: Report not found for the given sale_id
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       500:
     *         description: Internal server error
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     */
    app.use('/report', (req, res) => {
        proxyRequest(process.env.REPORT_SERVICE_URL as string, 'report', req, res)
    })

    /**
     * @openapi
     * /report/api/v1/reports/{user_id}/purchases:
     *   get:
     *     tags:
     *       - Report
     *     summary: Request a sale by user ID
     *     description: This endpoint proxies requests to the report service to retrieve a report by the specified sale ID.
     *     parameters:
     *       - in: path
     *         name: user_id
     *         schema:
     *           type: string
     *         description: Unique identifier for the sale to retrieve the report for
     *         required: true
     *     responses:
     *       200:
     *         description: Successfully retrieved the report for the given sale ID
     *         content:
     *           application/json:
     *             schema:
     *               type: object
     *               properties:
     *                 UserID:
     *                   type: string
     *                   description: Unique identifier for the user who made the sale
     *                 SaleID:
     *                   type: string
     *                   description: Unique identifier for the sale
     *                 PaymentType:
     *                   type: string
     *                   description: Payment method used for the sale
     *                 date:
     *                   type: string
     *                   format: date-time
     *                   description: Date and time when the sale occurred
     *                 Total:
     *                   type: integer
     *                   description: Total amount of the sale
     *                 Products:
     *                   type: array
     *                   items:
     *                     type: object
     *                     properties:
     *                       ProductID:
     *                         type: string
     *                         description: Unique identifier for the product
     *                       Name:
     *                         type: string
     *                         description: Name of the product
     *                       Unit:
     *                         type: string
     *                         description: Unit of measurement for the product
     *                       Category:
     *                         type: string
     *                         description: Category of the product
     *                       Quantity:
     *                         type: integer
     *                         description: Quantity of the product sold
     *                       Price:
     *                         type: integer
     *                         description: Price of the product
     *             examples:
     *               application/json:
     *                 summary: Sample response
     *                 value:
     *                   UserID: "01J1P83E4KJ3EM3H73139AC5PQ"
     *                   SaleID: "01J3RN9RAGGGM1WWVHY2P5BEDW"
     *                   PaymentType: "credit"
     *                   date: "2024-07-26T20:16:30.160995Z"
     *                   Total: 62
     *                   Products:
     *                     - ProductID: "01H7D8E4M7TX2R3ZQ8G5V2C3Y4"
     *                       Name: "Refrigerante Coca Cola Lata 350ml"
     *                       Unit: "un"
     *                       Category: "Bebidas"
     *                       Quantity: 2
     *                       Price: 8
     *       400:
     *         description: Bad request due to missing or invalid sale_id
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       404:
     *         description: Report not found for the given sale_id
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     *       500:
     *         description: Internal server error
     *         content:
     *           application/json:
     *             schema:
     *               $ref: '#/components/schemas/ErrorResponse'
     */
    app.use('/report', (req, res) => {
        proxyRequest(process.env.REPORT_SERVICE_URL as string, 'report', req, res)
    })
}
