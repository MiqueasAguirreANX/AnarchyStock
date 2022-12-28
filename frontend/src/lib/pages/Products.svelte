<script>
    import {
        DataTable,
        Toolbar,
        ToolbarContent,
        ToolbarSearch,
        Modal,
        Button,
        Form,
        TextInput,
        NumberInput,
        PasswordInput,
        DataTableSkeleton,
        Pagination,
    } from "carbon-components-svelte";
    import { Grid, Row, Column } from "carbon-components-svelte";
    import Add from "carbon-icons-svelte/lib/Add.svelte";
    import Renew from "carbon-icons-svelte/lib/Renew.svelte";
    import Search from "carbon-icons-svelte/lib/Search.svelte";
    import { InlineLoading } from "carbon-components-svelte";

    let pageSize = 10;
    let page = 1;
    let filteredRowIds = [];
    let ref = null
    let loading = true;
    async function fetchData () {
        loading = true
		const response = await fetch('http://127.0.0.1:8080/products')
        const jsonResponse = await response.json()
        let rows = jsonResponse.message.map((elem)=>{
            elem.id = elem.ID
            return elem
        })
        loading = false
        return rows
	}
    let rowsPromise = fetchData()

    const tryAgain = ()=>{
        rowsPromise = fetchData()
    }
    async function addProduct(ev) {
        const formData = new FormData(ref);
        const data = {};
        for (let field of formData) {
            const [key, value] = field;
            data[key] = value;
        }
        data["Price"] = parseFloat(data["Price"])
        data["Quantity"] = parseInt(data["Quantity"])

        const response = await fetch('http://127.0.0.1:8080/products/create', {
            method: "POST",
            body: JSON.stringify(data),
        })
        const jsonResponse = await response.json()
        openAdd = false;
    }

    let openAdd = false;
    let openSearch = false;

    async function searchProduct(ev) {
        
    }
</script>
  
<div class="page">
    <Grid>
        <Row>
            <Column>
                <Button 
                    kind="ghost" 
                    iconDescription="Reload" 
                    icon={Renew}
                    on:click={(ev)=>{
                        tryAgain()
                    }}
                />
            </Column>
            <Column>
                <Button 
                    kind="ghost" 
                    iconDescription="Search" 
                    icon={Search}
                    on:click={()=> openSearch = true}
                />
            </Column>
            <Column>
                <Button 
                    kind="ghost" 
                    iconDescription="Add product" 
                    icon={Add}
                    on:click={()=>(openAdd = true)}
                />
            </Column>
        </Row>
    </Grid>
    
    <br>
    <br>

{#await rowsPromise}
    <InlineLoading></InlineLoading>
    <DataTableSkeleton showHeader={false} showToolbar={false} />
{:then rows}
    <DataTable
        stickyHeader
        headers={[
            { key: "ID",        value: "ID"},
            { key: "name",      value: "Name" },
            { key: "category",  value: "Category" },
            { key: "price",     value: "Price" },
            { key: "quantity",  value: "Quantity" },
        ]}
        {rows}
        {page}
        {pageSize}
    >
        <!-- <Toolbar>
            <ToolbarContent>
                <ToolbarSearch
                    persistent
                    value="round"
                    shouldFilterRows
                    bind:filteredRowIds
                />
            </ToolbarContent>
        </Toolbar> -->
    </DataTable>
    <Pagination 
        bind:pageSize
        bind:page
        totalItems={rows.length}
        pageSizeInputDisabled
    />
{:catch error}
    <p>Could not fetch data.</p>
    <p style="color: red">{error.message}</p>
{/await}



    <Modal
        bind:open={openAdd}
        modalHeading="Add product"
        primaryButtonText="Create"
        secondaryButtonText="Cancel"
        on:click:button--secondary={() => (openAdd = false)}
        on:open
        on:close
        on:submit={addProduct}
    >  
        <Form
            on:submit={(e) => {
                e.preventDefault();
            }}
            bind:ref={ref}
        >
            <TextInput labelText="Name" name="Name" placeholder="Enter Name..." required />
            <br>
            <TextInput labelText="Category" name="Category" placeholder="ex. Food" required />
            <br>
            <NumberInput label="Price" name="Price" placeholder="ex. 12.54" min={0} required />
            <br>
            <NumberInput label="Quantity" name="Quantity" placeholder="ex. 150" min={1} required />
            <br>
        </Form>
    </Modal>
    <Modal
        bind:open={openSearch}
        modalHeading="Search product"
        passiveModal
        on:click:button--secondary={() => (openSearch = false)}
        on:open
        on:close
        on:submit={searchProduct}
    >  
        <TextInput labelText="Name" name="Name" placeholder="Enter Name..." required />
        <br>
        <TextInput labelText="Category" name="Category" placeholder="ex. Food" required />
        <br>
        Search
    </Modal>
</div>


<style>
.page {
    width: 100%;
    height: 100%;
}
</style>