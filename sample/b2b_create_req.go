package sample

const B2bCreateWorkItemWith2Parties = `
<SOAP:Envelope xmlns:SOAP="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:schemas-xtrac-fmr-com:b2b" SOAP:encodingStyle="urn:encodings-xtrac-fmr-com:v2.0">
	<SOAP:Header>
		<Cookie>123456789012345678901234567890</Cookie>
		<Validate>1</Validate>
	</SOAP:Header>
	<SOAP:Body>
		<CreateWorkItem>
			<CurrentNode>RSCAN</CurrentNode>
		    <ItemType>DNACCT</ItemType>
		    <ItemSubtype>DUSA</ItemSubtype>
		    <Description>Create a UD Workitem through XAVI</Description>
			<Status>DOPEN</Status>
			<Cause/>
			<CommType>DMAIL</CommType>
			<Memo>some memo</Memo>
			<Amount>10</Amount>
			<ResolutionNote>
				<Name>resolution note name</Name>
				<Memo>some memo 2</Memo>
			</ResolutionNote>
			<TSFs>
				<TSF Name="AB#"/>
				<TSF Name="RRQT">FALSE</TSF>
				<TSF Name="ENV#"/>
				<TSF Name="ARCH"/>
			</TSFs>
			<Parties>
				<Party Name="ORIG">
				   <AccountNo>NoA/P-3</AccountNo>
				   <AccountType>Retail</AccountType>
				   <ActiveParty>Y</ActiveParty>
				   <OrgName>NoA/P-3</OrgName>
				   <PartyIdType>Other</PartyIdType>
				   <PartyId/>
				</Party>
				<Party Name="DSPS">
				   <AccountNo>01234560</AccountNo>
				   <AccountType>Retail</AccountType>
				   <ActiveParty>Y</ActiveParty>
				   <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
				 </Party>
			 </Parties>
      	</CreateWorkItem>
   </SOAP:Body>
</SOAP:Envelope>
`

const B2bCreateWorkItemWith20Parties = `
<SOAP:Envelope xmlns:SOAP="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:schemas-xtrac-fmr-com:b2b" SOAP:encodingStyle="urn:encodings-xtrac-fmr-com:v2.0">
	<SOAP:Header>
		<Cookie>123456789012345678901234567890</Cookie>
		<Validate>1</Validate>
	</SOAP:Header>
	<SOAP:Body>
		<CreateWorkItem>
			<CurrentNode>RSCAN</CurrentNode>
		    <ItemType>DNACCT</ItemType>
		    <ItemSubtype>DUSA</ItemSubtype>
		    <Description>Create a UD Workitem through XAVI</Description>
		        <Status>DOPEN</Status>
		        <Cause/>
		        <CommType>DMAIL</CommType>
					<Memo>some memo</Memo>
					<Amount>10</Amount>
				<ResolutionNote>
					<Name>resolution note name</Name>
					<Memo>some memo 2</Memo>
				</ResolutionNote>
		        <TSFs>
					<TSF Name="AB#"/>
		            <TSF Name="RRQT">FALSE</TSF>
		            <TSF Name="ENV#"/>
		            <TSF Name="ARCH"/>
		        </TSFs>
		        <Parties>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		         </Parties>
      </CreateWorkItem>
   </SOAP:Body>
</SOAP:Envelope>
`

const B2bCreateWorkItemWith200Parties = `
<SOAP:Envelope xmlns:SOAP="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:schemas-xtrac-fmr-com:b2b" SOAP:encodingStyle="urn:encodings-xtrac-fmr-com:v2.0">
	<SOAP:Header>
		<Cookie>123456789012345678901234567890</Cookie>
		<Validate>1</Validate>
	</SOAP:Header>
	<SOAP:Body>
		<CreateWorkItem>
			<CurrentNode>RSCAN</CurrentNode>
		    <ItemType>DNACCT</ItemType>
		    <ItemSubtype>DUSA</ItemSubtype>
		    <Description>Create a UD Workitem through XAVI</Description>
		        <Status>DOPEN</Status>
		        <Cause/>
		        <CommType>DMAIL</CommType>
					<Memo>some memo</Memo>
					<Amount>10</Amount>
				<ResolutionNote>
					<Name>resolution note name</Name>
					<Memo>some memo 2</Memo>
				</ResolutionNote>
		        <TSFs>
					<TSF Name="AB#"/>
		            <TSF Name="RRQT">FALSE</TSF>
		            <TSF Name="ENV#"/>
		            <TSF Name="ARCH"/>
		        </TSFs>
		        <Parties>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		         </Parties>
      </CreateWorkItem>
   </SOAP:Body>
</SOAP:Envelope>
`

const B2bCreateWorkItemWith2000Parties = `
<SOAP:Envelope xmlns:SOAP="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:schemas-xtrac-fmr-com:b2b" SOAP:encodingStyle="urn:encodings-xtrac-fmr-com:v2.0">
	<SOAP:Header>
		<Cookie>123456789012345678901234567890</Cookie>
		<Validate>1</Validate>
	</SOAP:Header>
	<SOAP:Body>
		<CreateWorkItem>
			<CurrentNode>RSCAN</CurrentNode>
		    <ItemType>DNACCT</ItemType>
		    <ItemSubtype>DUSA</ItemSubtype>
		    <Description>Create a UD Workitem through XAVI</Description>
		        <Status>DOPEN</Status>
		        <Cause/>
		        <CommType>DMAIL</CommType>
					<Memo>some memo</Memo>
					<Amount>10</Amount>
				<ResolutionNote>
					<Name>resolution note name</Name>
					<Memo>some memo 2</Memo>
				</ResolutionNote>
		        <TSFs>
					<TSF Name="AB#"/>
		            <TSF Name="RRQT">FALSE</TSF>
		            <TSF Name="ENV#"/>
		            <TSF Name="ARCH"/>
		        </TSFs>
		        <Parties>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		            <Party Name="ORIG">
		               <AccountNo>NoA/P-3</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>NoA/P-3</OrgName>
		               <PartyIdType>Other</PartyIdType>
		               <PartyId/>
		            </Party>
		            <Party Name="DSPS">
		               <AccountNo>01234560</AccountNo>
		               <AccountType>Retail</AccountType>
		               <ActiveParty>Y</ActiveParty>
		               <OrgName>DOPEN/010698/10/RRQT False2of3</OrgName>
		             </Party>
		         </Parties>
      </CreateWorkItem>
   </SOAP:Body>
</SOAP:Envelope>
`
